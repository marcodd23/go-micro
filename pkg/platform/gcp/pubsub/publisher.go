package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/marcodd23/go-micro/pkg/messaging/publisher"
	"google.golang.org/api/option"
	"time"
)

// NewPubSubBufferedPublisherFactory - factory that create a pubsub client and then initialize a publisher.BufferedPublisher.
func NewPubSubBufferedPublisherFactory(
	ctx context.Context,
	projectID string,
	batchSize int32,
	flushDelayThreshold *time.Duration,
	opts ...option.ClientOption) (publisher.BufferedPublisher, error) {
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, publisher.NewMessagingErrorCode(publisher.ErrorInitializingPubsubClient, err)
	}

	publishConfig := publisher.TopicPublishConfig{
		BatchSize: batchSize,
	}

	if flushDelayThreshold == nil {
		publishConfig.FlushDelayThreshold = publisher.DefaultFlushDelayThreshold
	} else {
		publishConfig.FlushDelayThreshold = *flushDelayThreshold
	}

	pubSubClient := &pubSubClient{client: client}

	return publisher.NewBufferedPublisher(pubSubClient, publishConfig)
}

// NewBufferedPublisherWithRetryFactory - factory that create a pubsub client and then initialize a publisher.BufferedPublisherWithRetry.
func NewBufferedPublisherWithRetryFactory(
	ctx context.Context,
	projectID string,
	batchSize int32,
	flushDelayThreshold *time.Duration,
	maxRetryCount int16,
	initialRetryInterval *time.Duration,
	opts ...option.ClientOption) (publisher.BufferedPublisherWithRetry, error) {
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, publisher.NewMessagingErrorCode(publisher.ErrorInitializingPubsubClient, err)
	}

	publishConfig := publisher.TopicPublishConfig{
		BatchSize: batchSize,
	}

	if flushDelayThreshold == nil {
		publishConfig.FlushDelayThreshold = publisher.DefaultFlushDelayThreshold
	} else {
		publishConfig.FlushDelayThreshold = *flushDelayThreshold
	}

	if initialRetryInterval == nil {
		publishConfig.InitialRetryInterval = publisher.DefaultInitialRetryInterval
	} else {
		publishConfig.InitialRetryInterval = *initialRetryInterval
	}

	pubSubClient := &pubSubClient{client: client}

	return publisher.NewBufferedPublisherWithRetry(ctx, pubSubClient, publishConfig)
}
