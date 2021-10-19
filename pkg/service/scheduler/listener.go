// Copyright (C) 2017 ScyllaDB

package scheduler

import (
	"context"
	"time"

	"github.com/scylladb/go-log"
	"github.com/scylladb/scylla-manager/pkg/scheduler"
)

type schedulerListener struct {
	logger log.Logger
}

func (l schedulerListener) OnSchedulerStart(ctx context.Context) {
	l.logger.Debug(ctx, "OnSchedulerStart")
}

func (l schedulerListener) OnSchedulerStop(ctx context.Context) {
	l.logger.Debug(ctx, "OnSchedulerStop")
}

func (l schedulerListener) OnRunStart(ctx *scheduler.RunContext) {
	l.logger.Debug(ctx, "OnRunStart", "key", ctx.Key, "retry", ctx.Retry)
}

func (l schedulerListener) OnRunSuccess(ctx *scheduler.RunContext) {
	l.logger.Debug(ctx, "OnRunSuccess", "key", ctx.Key, "retry", ctx.Retry)
}

func (l schedulerListener) OnRunStop(ctx *scheduler.RunContext) {
	l.logger.Debug(ctx, "OnRunStop", "key", ctx.Key, "retry", ctx.Retry)
}

func (l schedulerListener) OnRunWindowEnd(ctx *scheduler.RunContext) {
	l.logger.Debug(ctx, "OnRunWindowEnd", "key", ctx.Key, "retry", ctx.Retry)
}

func (l schedulerListener) OnRunError(ctx *scheduler.RunContext, err error) {
	l.logger.Debug(ctx, "OnRunError", "key", ctx.Key, "retry", ctx.Retry, "error", err)
}

func (l schedulerListener) OnSchedule(ctx context.Context, key scheduler.Key, begin, end time.Time, retno int8) {
	l.logger.Debug(ctx, "OnSchedule", "key", key, "begin", begin, "end", end, "retry", retno)
}

func (l schedulerListener) OnUnschedule(ctx context.Context, key scheduler.Key) {
	l.logger.Debug(ctx, "OnUnschedule", "key", key)
}

func (l schedulerListener) OnTrigger(ctx context.Context, key scheduler.Key, success bool) {
	l.logger.Debug(ctx, "OnTrigger", "key", key, "success", success)
}

func (l schedulerListener) OnStop(ctx context.Context, key scheduler.Key) {
	l.logger.Debug(ctx, "OnStop", "key", key)
}

func (l schedulerListener) OnRetryBackoff(ctx context.Context, key scheduler.Key, backoff time.Duration, retno int8) {
	l.logger.Debug(ctx, "OnRetryBackoff", "key", key, "backoff", retno)
}

func (l schedulerListener) OnNoTrigger(ctx context.Context, key scheduler.Key) {
	l.logger.Debug(ctx, "OnNoTrigger", "key", key)
}

func (l schedulerListener) OnSleep(ctx context.Context, key scheduler.Key, d time.Duration) {
	l.logger.Debug(ctx, "OnSleep", "key", key, "duration", d)
}
