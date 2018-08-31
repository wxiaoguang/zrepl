package job

import (
	"context"
	"github.com/pkg/errors"
	"github.com/problame/go-streamrpc"
	"github.com/zrepl/zrepl/config"
	"github.com/zrepl/zrepl/daemon/connecter"
	"github.com/zrepl/zrepl/daemon/filters"
	"github.com/zrepl/zrepl/daemon/pruner"
	"github.com/zrepl/zrepl/endpoint"
	"github.com/zrepl/zrepl/replication"
	"sync"
	"github.com/zrepl/zrepl/daemon/logging"
)

type Push struct {
	name      string
	connecter streamrpc.Connecter
	fsfilter  endpoint.FSFilter

	prunerFactory *pruner.PrunerFactory

	mtx         sync.Mutex
	replication *replication.Replication
}

func PushFromConfig(g *config.Global, in *config.PushJob) (j *Push, err error) {

	j = &Push{}
	j.name = in.Name

	j.connecter, err = connecter.FromConfig(g, in.Replication.Connect)

	if j.fsfilter, err = filters.DatasetMapFilterFromConfig(in.Replication.Filesystems); err != nil {
		return nil, errors.Wrap(err, "cannnot build filesystem filter")
	}

	j.prunerFactory, err = pruner.NewPrunerFactory(in.Pruning)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (j *Push) Name() string { return j.name }

func (j *Push) Status() interface{} {
	rep := func() *replication.Replication {
		j.mtx.Lock()
		defer j.mtx.Unlock()
		if j.replication == nil {
			return nil
		}
		return j.replication
	}()
	if rep == nil {
		return nil
	}
	return rep.Report()
}

func (j *Push) Run(ctx context.Context) {
	log := GetLogger(ctx)

	defer log.Info("job exiting")

	invocationCount := 0
outer:
	for {
		log.Info("wait for wakeups")
		select {
		case <-ctx.Done():
			log.WithError(ctx.Err()).Info("context")
			break outer
		case <-WaitWakeup(ctx):
			invocationCount++
			invLog := log.WithField("invocation", invocationCount)
			j.do(WithLogger(ctx, invLog))
		}
	}
}

func (j *Push) do(ctx context.Context) {

	log := GetLogger(ctx)
	ctx = logging.WithSubsystemLoggers(ctx, log)

	client, err := streamrpc.NewClient(j.connecter, &streamrpc.ClientConfig{STREAMRPC_CONFIG})
	if err != nil {
		log.WithError(err).Error("cannot create streamrpc client")
	}
	defer client.Close(ctx)

	sender := endpoint.NewSender(j.fsfilter, filters.NewAnyFSVFilter())
	receiver := endpoint.NewRemote(client)

	j.mtx.Lock()
	j.replication = replication.NewReplication()
	j.mtx.Unlock()

	log.Info("start replication")
	j.replication.Drive(ctx, sender, receiver)

	log.Info("start pruning sender")
	psCtx := pruner.WithLogger(ctx, pruner.GetLogger(ctx).WithField("prune_side", "sender"))
	senderPruner := j.prunerFactory.BuildSenderPruner(psCtx, sender, sender) // FIXME ctx as member
	senderPruner.Prune()

	log.Info("start pruning receiver")
	prCtx := pruner.WithLogger(ctx, pruner.GetLogger(ctx).WithField("prune_side", "receiver"))
	receiverPruner := j.prunerFactory.BuildReceiverPruner(prCtx, receiver, sender) // FIXME ctx as member
	receiverPruner.Prune()

}
