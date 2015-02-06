package client

const (
	VOLUME_TYPE = "volume"
)

type Volume struct {
	Resource
    
    AccountId string `json:"accountId,omitempty"`
    
    Created string `json:"created,omitempty"`
    
    Data map[string]interface{} `json:"data,omitempty"`
    
    Description string `json:"description,omitempty"`
    
    ImageId string `json:"imageId,omitempty"`
    
    InstanceId string `json:"instanceId,omitempty"`
    
    IsHostPath bool `json:"isHostPath,omitempty"`
    
    Kind string `json:"kind,omitempty"`
    
    Name string `json:"name,omitempty"`
    
    RemoveTime string `json:"removeTime,omitempty"`
    
    Removed string `json:"removed,omitempty"`
    
    State string `json:"state,omitempty"`
    
    Transitioning string `json:"transitioning,omitempty"`
    
    TransitioningMessage string `json:"transitioningMessage,omitempty"`
    
    TransitioningProgress int `json:"transitioningProgress,omitempty"`
    
    Uri string `json:"uri,omitempty"`
    
    Uuid string `json:"uuid,omitempty"`
    
}

type VolumeCollection struct {
	Collection
	Data []Volume `json:"data,omitempty"`
}

type VolumeClient struct {
	rancherClient *RancherClient
}

type VolumeOperations interface {
	List(opts *ListOpts) (*VolumeCollection, error)
	Create(opts *Volume) (*Volume, error)
	Update(existing *Volume, updates interface{}) (*Volume, error)
	ById(id string) (*Volume, error)
	Delete(container *Volume) error
}

func newVolumeClient(rancherClient *RancherClient) *VolumeClient {
	return &VolumeClient{
		rancherClient: rancherClient,
	}
}

func (c *VolumeClient) Create(container *Volume) (*Volume, error) {
	resp := &Volume{}
	err := c.rancherClient.doCreate(VOLUME_TYPE, container, resp)
	return resp, err
}

func (c *VolumeClient) Update(existing *Volume, updates interface{}) (*Volume, error) {
	resp := &Volume{}
	err := c.rancherClient.doUpdate(VOLUME_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VolumeClient) List(opts *ListOpts) (*VolumeCollection, error) {
	resp := &VolumeCollection{}
	err := c.rancherClient.doList(VOLUME_TYPE, opts, resp)
	return resp, err
}

func (c *VolumeClient) ById(id string) (*Volume, error) {
	resp := &Volume{}
	err := c.rancherClient.doById(VOLUME_TYPE, id, resp)
	return resp, err
}

func (c *VolumeClient) Delete(container *Volume) error {
	return c.rancherClient.doResourceDelete(VOLUME_TYPE, &container.Resource)
}
