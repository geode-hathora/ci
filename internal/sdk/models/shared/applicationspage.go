// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ApplicationsPage struct {
	Applications []ApplicationWithLatestDeploymentAndBuild `json:"applications"`
}

func (o *ApplicationsPage) GetApplications() []ApplicationWithLatestDeploymentAndBuild {
	if o == nil {
		return []ApplicationWithLatestDeploymentAndBuild{}
	}
	return o.Applications
}
