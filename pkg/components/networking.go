package components

import (
	"context"
	"github.com/openshift/microshift/pkg/assets"
	"github.com/openshift/microshift/pkg/config"
	"k8s.io/klog/v2"
)

func startCNIPlugin(ctx context.Context, _ *config.Config, kubeconfigPath string) error {
	var (
		ns = []string{
			"components/flannel/namespace.yaml",
		}
		sa = []string{
			"components/flannel/sa.yaml",
		}
		cr = []string{
			"components/flannel/clusterrole.yaml",
		}
		crb = []string{
			"components/flannel/clusterrolebinding.yaml",
		}
		cm = []string{
			"components/flannel/configmap.yaml",
		}
		apps = []string{
			"components/flannel/daemonset.yaml",
		}
	)

	if err := assets.ApplyNamespaces(ctx, ns, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply ns %v: %v", ns, err)
		return err
	}
	if err := assets.ApplyServiceAccounts(ctx, sa, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply serviceAccount %v %v", sa, err)
		return err
	}
	if err := assets.ApplyClusterRoles(ctx, cr, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply clusterRole %v %v", cr, err)
		return err
	}
	if err := assets.ApplyClusterRoleBindings(ctx, crb, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply clusterRoleBinding %v %v", crb, err)
		return err
	}
	if err := assets.ApplyConfigMaps(ctx, cm, nil, map[string]interface{}{}, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply configMap %v %v", cm, err)
		return err
	}
	if err := assets.ApplyDaemonSets(ctx, apps, nil, map[string]interface{}{}, kubeconfigPath); err != nil {
		klog.Warningf("Failed to apply apps %v %v", apps, err)
		return err
	}
	return nil
}
