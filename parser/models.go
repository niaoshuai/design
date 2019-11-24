package parser


type DConfig struct {
	Key   string
	Value string
}

type DComponent struct {
	Name            string `json:"name"`
	ChildComponents []DComponent
	Configs         map[string]string
}

type DSee struct {
	ComponentName string
	Data          string
}

type DDo struct {
	UIEvent       string
	ComponentName string
	Data          string
}

type DReact struct {
	SceneName          string
	ReactAction        string
	ReactComponentName string
	AnimateName        string
	ReactComponentData string
}

type DInteraction struct {
	See   DSee
	Do    DDo
	React []DReact
}

type DFlow struct {
	Interactions []DInteraction
	FlowName     string
}

func CreateDComponent(componentName string) *DComponent {
	return &DComponent{componentName, nil, nil}
}


type DLayout struct {
	LayoutName string
	LayoutRows []DLayoutRow
}

type DLayoutRow struct {
	ComponentName string
	LayoutInformation string
	NormalInformation []string
}

func CreateInteraction() *DInteraction {
	seeModel := &DSee{"", ""}
	doModel := &DDo{"", "", ""}
	return &DInteraction{
		See:   *seeModel,
		Do:    *doModel,
		React: nil,
	}
}
