package deployment

import (
	"context"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	autoscalingv1beta1 "github.com/mercari/tortoise/api/v1beta1"
)

type Service struct {
	c client.Client
}

func New(c client.Client) *Service {
	return &Service{c: c}
}

func (c *Service) GetDeploymentOnTortoise(ctx context.Context, tortoise *autoscalingv1beta1.Tortoise) (*v1.Deployment, error) {
	d := &v1.Deployment{}
	if err := c.c.Get(ctx, types.NamespacedName{Namespace: tortoise.Namespace, Name: tortoise.Spec.TargetRefs.ScaleTargetRef.Name}, d); err != nil {
		return nil, fmt.Errorf("failed to get deployment on tortoise: %w", err)
	}
	return d, nil
}
