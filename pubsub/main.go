package main

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	project = "your-gcp-project-id"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		topic, err := pubsub.NewTopic(ctx, "myTopic", &pubsub.TopicArgs{
			Project: pulumi.String(project),
		})
		if err != nil {
			return err
		}

		subscription, err := pubsub.NewSubscription(ctx, "mySubscription", &pubsub.SubscriptionArgs{
			Project: pulumi.String(project),
			Topic:   topic.Name,
		})
		if err != nil {
			return err
		}

		if err := ctx.RegisterResource("gcp:pubsub/topic:Topic", topic); err != nil {
			return err
		}
		if err := ctx.RegisterResource("gcp:pubsub/subscription:Subscription", subscription); err != nil {
			return err
		}

		publisher := func() error {
			message := "Hello, World!"
			return publishMessage(ctx, topic.Name, message)
		}

		consumer := func() error {
			return consumeMessages(ctx, subscription.Name)
		}

		go func() {
			if err := publisher(); err != nil {
				log.Fatalf("Failed to publish message: %v", err)
			}
		}()

		if err := consumer(); err != nil {
			log.Fatalf("Failed to consume messages: %v", err)
		}

		return nil
	})
}

func publishMessage(ctx *pulumi.Context, topicName pulumi.StringOutput, message string) error {
	topicRef := topicName.ApplyT(func(name string) (string, error) {
		return fmt.Sprintf("projects/%s/topics/%s", project, name), nil
	}).(pulumi.StringOutput)

	pubsubMessage := pubsub.MessageArgs{
		Data: pulumi.String(message),
	}

	_, err := pubsub.NewMessage(ctx, "myMessage", &pubsubMessage, pulumi.DependsOn([]pulumi.Resource{topicRef}))
	return err
}

func consumeMessages(ctx *pulumi.Context, subscriptionName pulumi.StringOutput) error {
	subscriptionRef := subscriptionName.ApplyT(func(name string) (string, error) {
		return fmt.Sprintf("projects/%s/subscriptions/%s", project, name), nil
	}).(pulumi.StringOutput)

	trigger := pubsub.NewSubscriptionIAMMember(ctx, "myTrigger", &pubsub.SubscriptionIAMMemberArgs{
		Project:       pulumi.String(project),
		Subscription:  subscriptionRef,
		Role:          pulumi.String("roles/pubsub.subscriber"),
		Member:        pulumi.String("allUsers"),
		ManagePolicy:  pulumi.Bool(true),
		ManageMembers: pulumi.Bool(true),
	})

	if err := ctx.RegisterResource("gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember", trigger); err != nil {
		return err
	}

	err := pulumi.All(trigger.ID()).ApplyT(func(args interface{}) error {
		log.Printf("Subscription trigger completed.")
		return nil
	}).(pulumi.ResourceTransformation)

	return err
}
