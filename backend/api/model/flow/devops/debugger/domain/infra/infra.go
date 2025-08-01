// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package infra

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// ComponentType 支持组件类型
type ComponentType int64

const (
	ComponentType_Undefined ComponentType = 0
	// Coze Plugin
	ComponentType_CozePlugin ComponentType = 10000
	// Coze Tool
	ComponentType_CozeTool ComponentType = 10001
	// Coze Workflow
	ComponentType_CozeWorkflow ComponentType = 10002
	// Coze SubWorkflow，即在Workflow中被引用的子Workflow
	ComponentType_CozeSubWorkflow ComponentType = 10003
	// Coze workflow中的LLM节点
	ComponentType_CozeLLMNode ComponentType = 10004
	// Coze workflow中的Code节点
	ComponentType_CozeCodeNode ComponentType = 10005
	// Coze workflow中的Knowledge节点
	ComponentType_CozeKnowledgeNode ComponentType = 10006
	// Coze workflow中的Tool节点
	ComponentType_CozeToolNode ComponentType = 10007
	// Coze workflow中的start节点
	ComponentType_CozeStartNode ComponentType = 10008
	// Coze workflow中的variable节点
	ComponentType_CozeVariableNode ComponentType = 10009
	// Coze 虚拟节点用于标识 variable 依赖的bot
	ComponentType_CozeVariableBot ComponentType = 20000
	// Coze 虚拟节点用于标识 variable 依赖的chat
	ComponentType_CozeVariableChat ComponentType = 20001
)

func (p ComponentType) String() string {
	switch p {
	case ComponentType_Undefined:
		return "Undefined"
	case ComponentType_CozePlugin:
		return "CozePlugin"
	case ComponentType_CozeTool:
		return "CozeTool"
	case ComponentType_CozeWorkflow:
		return "CozeWorkflow"
	case ComponentType_CozeSubWorkflow:
		return "CozeSubWorkflow"
	case ComponentType_CozeLLMNode:
		return "CozeLLMNode"
	case ComponentType_CozeCodeNode:
		return "CozeCodeNode"
	case ComponentType_CozeKnowledgeNode:
		return "CozeKnowledgeNode"
	case ComponentType_CozeToolNode:
		return "CozeToolNode"
	case ComponentType_CozeStartNode:
		return "CozeStartNode"
	case ComponentType_CozeVariableNode:
		return "CozeVariableNode"
	case ComponentType_CozeVariableBot:
		return "CozeVariableBot"
	case ComponentType_CozeVariableChat:
		return "CozeVariableChat"
	}
	return "<UNSET>"
}

func ComponentTypeFromString(s string) (ComponentType, error) {
	switch s {
	case "Undefined":
		return ComponentType_Undefined, nil
	case "CozePlugin":
		return ComponentType_CozePlugin, nil
	case "CozeTool":
		return ComponentType_CozeTool, nil
	case "CozeWorkflow":
		return ComponentType_CozeWorkflow, nil
	case "CozeSubWorkflow":
		return ComponentType_CozeSubWorkflow, nil
	case "CozeLLMNode":
		return ComponentType_CozeLLMNode, nil
	case "CozeCodeNode":
		return ComponentType_CozeCodeNode, nil
	case "CozeKnowledgeNode":
		return ComponentType_CozeKnowledgeNode, nil
	case "CozeToolNode":
		return ComponentType_CozeToolNode, nil
	case "CozeStartNode":
		return ComponentType_CozeStartNode, nil
	case "CozeVariableNode":
		return ComponentType_CozeVariableNode, nil
	case "CozeVariableBot":
		return ComponentType_CozeVariableBot, nil
	case "CozeVariableChat":
		return ComponentType_CozeVariableChat, nil
	}
	return ComponentType(0), fmt.Errorf("not a valid ComponentType string")
}

func ComponentTypePtr(v ComponentType) *ComponentType { return &v }
func (p *ComponentType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ComponentType(result.Int64)
	return
}

func (p *ComponentType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// TrafficScene 流量请求场景
type TrafficScene int64

const (
	TrafficScene_Undefined TrafficScene = 0
	// 单Agent调试页
	TrafficScene_CozeSingleAgentDebug TrafficScene = 10000
	// 多Agent调试页
	TrafficScene_CozeMultiAgentDebug TrafficScene = 10001
	// Tool调试页
	TrafficScene_CozeToolDebug TrafficScene = 10002
	// Workflow调试页
	TrafficScene_CozeWorkflowDebug TrafficScene = 10003
)

func (p TrafficScene) String() string {
	switch p {
	case TrafficScene_Undefined:
		return "Undefined"
	case TrafficScene_CozeSingleAgentDebug:
		return "CozeSingleAgentDebug"
	case TrafficScene_CozeMultiAgentDebug:
		return "CozeMultiAgentDebug"
	case TrafficScene_CozeToolDebug:
		return "CozeToolDebug"
	case TrafficScene_CozeWorkflowDebug:
		return "CozeWorkflowDebug"
	}
	return "<UNSET>"
}

func TrafficSceneFromString(s string) (TrafficScene, error) {
	switch s {
	case "Undefined":
		return TrafficScene_Undefined, nil
	case "CozeSingleAgentDebug":
		return TrafficScene_CozeSingleAgentDebug, nil
	case "CozeMultiAgentDebug":
		return TrafficScene_CozeMultiAgentDebug, nil
	case "CozeToolDebug":
		return TrafficScene_CozeToolDebug, nil
	case "CozeWorkflowDebug":
		return TrafficScene_CozeWorkflowDebug, nil
	}
	return TrafficScene(0), fmt.Errorf("not a valid TrafficScene string")
}

func TrafficScenePtr(v TrafficScene) *TrafficScene { return &v }
func (p *TrafficScene) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = TrafficScene(result.Int64)
	return
}

func (p *TrafficScene) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// ComponentMappingType 组件映射类型
type ComponentMappingType int64

const (
	ComponentMappingType_Undefined ComponentMappingType = 0
	ComponentMappingType_MockSet   ComponentMappingType = 1
)

func (p ComponentMappingType) String() string {
	switch p {
	case ComponentMappingType_Undefined:
		return "Undefined"
	case ComponentMappingType_MockSet:
		return "MockSet"
	}
	return "<UNSET>"
}

func ComponentMappingTypeFromString(s string) (ComponentMappingType, error) {
	switch s {
	case "Undefined":
		return ComponentMappingType_Undefined, nil
	case "MockSet":
		return ComponentMappingType_MockSet, nil
	}
	return ComponentMappingType(0), fmt.Errorf("not a valid ComponentMappingType string")
}

func ComponentMappingTypePtr(v ComponentMappingType) *ComponentMappingType { return &v }
func (p *ComponentMappingType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ComponentMappingType(result.Int64)
	return
}

func (p *ComponentMappingType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type OrderBy int64

const (
	OrderBy_UpdateTime OrderBy = 1
)

func (p OrderBy) String() string {
	switch p {
	case OrderBy_UpdateTime:
		return "UpdateTime"
	}
	return "<UNSET>"
}

func OrderByFromString(s string) (OrderBy, error) {
	switch s {
	case "UpdateTime":
		return OrderBy_UpdateTime, nil
	}
	return OrderBy(0), fmt.Errorf("not a valid OrderBy string")
}

func OrderByPtr(v OrderBy) *OrderBy { return &v }
func (p *OrderBy) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = OrderBy(result.Int64)
	return
}

func (p *OrderBy) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type DebugScene int64

const (
	// 默认play ground Debug场景
	DebugScene_Debug DebugScene = 0
)

func (p DebugScene) String() string {
	switch p {
	case DebugScene_Debug:
		return "Debug"
	}
	return "<UNSET>"
}

func DebugSceneFromString(s string) (DebugScene, error) {
	switch s {
	case "Debug":
		return DebugScene_Debug, nil
	}
	return DebugScene(0), fmt.Errorf("not a valid DebugScene string")
}

func DebugScenePtr(v DebugScene) *DebugScene { return &v }
func (p *DebugScene) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = DebugScene(result.Int64)
	return
}

func (p *DebugScene) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type CozeChannel int64

const (
	// 默认为Coze, 未来扩展到其他渠道
	CozeChannel_Coze CozeChannel = 0
)

func (p CozeChannel) String() string {
	switch p {
	case CozeChannel_Coze:
		return "Coze"
	}
	return "<UNSET>"
}

func CozeChannelFromString(s string) (CozeChannel, error) {
	switch s {
	case "Coze":
		return CozeChannel_Coze, nil
	}
	return CozeChannel(0), fmt.Errorf("not a valid CozeChannel string")
}

func CozeChannelPtr(v CozeChannel) *CozeChannel { return &v }
func (p *CozeChannel) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = CozeChannel(result.Int64)
	return
}

func (p *CozeChannel) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// BizCtx 业务上下文
type BizCtx struct {
	// connectorID
	ConnectorID *string `thrift:"connectorID,1,optional" form:"connectorID" json:"connectorID,omitempty" query:"connectorID"`
	// connector下用户ID
	ConnectorUID *string `thrift:"connectorUID,2,optional" form:"connectorUID" json:"connectorUID,omitempty" query:"connectorUID"`
	// 业务场景
	TrafficScene *TrafficScene `thrift:"trafficScene,3,optional" form:"trafficScene" json:"trafficScene,omitempty" query:"trafficScene"`
	// 业务场景组件ID，比如Bot调试页，则trafficSceneID为BotID
	TrafficCallerID *string `thrift:"trafficCallerID,4,optional" form:"trafficCallerID" json:"trafficCallerID,omitempty" query:"trafficCallerID"`
	// 业务线SpaceID，用于访问控制
	BizSpaceID *string `thrift:"bizSpaceID,5,optional" form:"bizSpaceID" json:"bizSpaceID,omitempty" query:"bizSpaceID"`
	// 额外信息
	Ext map[string]string `thrift:"ext,6,optional" form:"ext" json:"ext,omitempty" query:"ext"`
}

func NewBizCtx() *BizCtx {
	return &BizCtx{}
}

func (p *BizCtx) InitDefault() {
}

var BizCtx_ConnectorID_DEFAULT string

func (p *BizCtx) GetConnectorID() (v string) {
	if !p.IsSetConnectorID() {
		return BizCtx_ConnectorID_DEFAULT
	}
	return *p.ConnectorID
}

var BizCtx_ConnectorUID_DEFAULT string

func (p *BizCtx) GetConnectorUID() (v string) {
	if !p.IsSetConnectorUID() {
		return BizCtx_ConnectorUID_DEFAULT
	}
	return *p.ConnectorUID
}

var BizCtx_TrafficScene_DEFAULT TrafficScene

func (p *BizCtx) GetTrafficScene() (v TrafficScene) {
	if !p.IsSetTrafficScene() {
		return BizCtx_TrafficScene_DEFAULT
	}
	return *p.TrafficScene
}

var BizCtx_TrafficCallerID_DEFAULT string

func (p *BizCtx) GetTrafficCallerID() (v string) {
	if !p.IsSetTrafficCallerID() {
		return BizCtx_TrafficCallerID_DEFAULT
	}
	return *p.TrafficCallerID
}

var BizCtx_BizSpaceID_DEFAULT string

func (p *BizCtx) GetBizSpaceID() (v string) {
	if !p.IsSetBizSpaceID() {
		return BizCtx_BizSpaceID_DEFAULT
	}
	return *p.BizSpaceID
}

var BizCtx_Ext_DEFAULT map[string]string

func (p *BizCtx) GetExt() (v map[string]string) {
	if !p.IsSetExt() {
		return BizCtx_Ext_DEFAULT
	}
	return p.Ext
}

var fieldIDToName_BizCtx = map[int16]string{
	1: "connectorID",
	2: "connectorUID",
	3: "trafficScene",
	4: "trafficCallerID",
	5: "bizSpaceID",
	6: "ext",
}

func (p *BizCtx) IsSetConnectorID() bool {
	return p.ConnectorID != nil
}

func (p *BizCtx) IsSetConnectorUID() bool {
	return p.ConnectorUID != nil
}

func (p *BizCtx) IsSetTrafficScene() bool {
	return p.TrafficScene != nil
}

func (p *BizCtx) IsSetTrafficCallerID() bool {
	return p.TrafficCallerID != nil
}

func (p *BizCtx) IsSetBizSpaceID() bool {
	return p.BizSpaceID != nil
}

func (p *BizCtx) IsSetExt() bool {
	return p.Ext != nil
}

func (p *BizCtx) Read(iprot thrift.TProtocol) (err error) {
	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.MAP {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BizCtx[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BizCtx) ReadField1(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ConnectorID = _field
	return nil
}
func (p *BizCtx) ReadField2(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ConnectorUID = _field
	return nil
}
func (p *BizCtx) ReadField3(iprot thrift.TProtocol) error {

	var _field *TrafficScene
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := TrafficScene(v)
		_field = &tmp
	}
	p.TrafficScene = _field
	return nil
}
func (p *BizCtx) ReadField4(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.TrafficCallerID = _field
	return nil
}
func (p *BizCtx) ReadField5(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.BizSpaceID = _field
	return nil
}
func (p *BizCtx) ReadField6(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return err
	}
	_field := make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_key = v
		}

		var _val string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_val = v
		}

		_field[_key] = _val
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return err
	}
	p.Ext = _field
	return nil
}

func (p *BizCtx) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("BizCtx"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *BizCtx) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetConnectorID() {
		if err = oprot.WriteFieldBegin("connectorID", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ConnectorID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *BizCtx) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetConnectorUID() {
		if err = oprot.WriteFieldBegin("connectorUID", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ConnectorUID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}
func (p *BizCtx) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTrafficScene() {
		if err = oprot.WriteFieldBegin("trafficScene", thrift.I32, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.TrafficScene)); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}
func (p *BizCtx) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTrafficCallerID() {
		if err = oprot.WriteFieldBegin("trafficCallerID", thrift.STRING, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.TrafficCallerID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}
func (p *BizCtx) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetBizSpaceID() {
		if err = oprot.WriteFieldBegin("bizSpaceID", thrift.STRING, 5); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.BizSpaceID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}
func (p *BizCtx) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetExt() {
		if err = oprot.WriteFieldBegin("ext", thrift.MAP, 6); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Ext)); err != nil {
			return err
		}
		for k, v := range p.Ext {
			if err := oprot.WriteString(k); err != nil {
				return err
			}
			if err := oprot.WriteString(v); err != nil {
				return err
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *BizCtx) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BizCtx(%+v)", *p)

}

// ComponentSubject 业务组件的二级结构
type ComponentSubject struct {
	// 组件ID，例如Tool ID、Node ID等
	ComponentID *string `thrift:"componentID,1,optional" form:"componentID" json:"componentID,omitempty" query:"componentID"`
	// 组件类型
	ComponentType *ComponentType `thrift:"componentType,2,optional" form:"componentType" json:"componentType,omitempty" query:"componentType"`
	// 父组件ID，例如Tool->Plugin, Node->Workflow
	ParentComponentID *string `thrift:"parentComponentID,3,optional" form:"parentComponentID" json:"parentComponentID,omitempty" query:"parentComponentID"`
	// 父组件类型
	ParentComponentType *ComponentType `thrift:"parentComponentType,4,optional" form:"parentComponentType" json:"parentComponentType,omitempty" query:"parentComponentType"`
}

func NewComponentSubject() *ComponentSubject {
	return &ComponentSubject{}
}

func (p *ComponentSubject) InitDefault() {
}

var ComponentSubject_ComponentID_DEFAULT string

func (p *ComponentSubject) GetComponentID() (v string) {
	if !p.IsSetComponentID() {
		return ComponentSubject_ComponentID_DEFAULT
	}
	return *p.ComponentID
}

var ComponentSubject_ComponentType_DEFAULT ComponentType

func (p *ComponentSubject) GetComponentType() (v ComponentType) {
	if !p.IsSetComponentType() {
		return ComponentSubject_ComponentType_DEFAULT
	}
	return *p.ComponentType
}

var ComponentSubject_ParentComponentID_DEFAULT string

func (p *ComponentSubject) GetParentComponentID() (v string) {
	if !p.IsSetParentComponentID() {
		return ComponentSubject_ParentComponentID_DEFAULT
	}
	return *p.ParentComponentID
}

var ComponentSubject_ParentComponentType_DEFAULT ComponentType

func (p *ComponentSubject) GetParentComponentType() (v ComponentType) {
	if !p.IsSetParentComponentType() {
		return ComponentSubject_ParentComponentType_DEFAULT
	}
	return *p.ParentComponentType
}

var fieldIDToName_ComponentSubject = map[int16]string{
	1: "componentID",
	2: "componentType",
	3: "parentComponentID",
	4: "parentComponentType",
}

func (p *ComponentSubject) IsSetComponentID() bool {
	return p.ComponentID != nil
}

func (p *ComponentSubject) IsSetComponentType() bool {
	return p.ComponentType != nil
}

func (p *ComponentSubject) IsSetParentComponentID() bool {
	return p.ParentComponentID != nil
}

func (p *ComponentSubject) IsSetParentComponentType() bool {
	return p.ParentComponentType != nil
}

func (p *ComponentSubject) Read(iprot thrift.TProtocol) (err error) {
	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ComponentSubject[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ComponentSubject) ReadField1(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ComponentID = _field
	return nil
}
func (p *ComponentSubject) ReadField2(iprot thrift.TProtocol) error {

	var _field *ComponentType
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := ComponentType(v)
		_field = &tmp
	}
	p.ComponentType = _field
	return nil
}
func (p *ComponentSubject) ReadField3(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ParentComponentID = _field
	return nil
}
func (p *ComponentSubject) ReadField4(iprot thrift.TProtocol) error {

	var _field *ComponentType
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		tmp := ComponentType(v)
		_field = &tmp
	}
	p.ParentComponentType = _field
	return nil
}

func (p *ComponentSubject) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ComponentSubject"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ComponentSubject) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetComponentID() {
		if err = oprot.WriteFieldBegin("componentID", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ComponentID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *ComponentSubject) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetComponentType() {
		if err = oprot.WriteFieldBegin("componentType", thrift.I32, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.ComponentType)); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}
func (p *ComponentSubject) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetParentComponentID() {
		if err = oprot.WriteFieldBegin("parentComponentID", thrift.STRING, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ParentComponentID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}
func (p *ComponentSubject) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetParentComponentType() {
		if err = oprot.WriteFieldBegin("parentComponentType", thrift.I32, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(int32(*p.ParentComponentType)); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ComponentSubject) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ComponentSubject(%+v)", *p)

}

type Creator struct {
	ID        *string `thrift:"ID,1,optional" form:"ID" json:"ID,omitempty" query:"ID"`
	Name      *string `thrift:"name,2,optional" form:"name" json:"name,omitempty" query:"name"`
	AvatarUrl *string `thrift:"avatarUrl,3,optional" form:"avatarUrl" json:"avatarUrl,omitempty" query:"avatarUrl"`
}

func NewCreator() *Creator {
	return &Creator{}
}

func (p *Creator) InitDefault() {
}

var Creator_ID_DEFAULT string

func (p *Creator) GetID() (v string) {
	if !p.IsSetID() {
		return Creator_ID_DEFAULT
	}
	return *p.ID
}

var Creator_Name_DEFAULT string

func (p *Creator) GetName() (v string) {
	if !p.IsSetName() {
		return Creator_Name_DEFAULT
	}
	return *p.Name
}

var Creator_AvatarUrl_DEFAULT string

func (p *Creator) GetAvatarUrl() (v string) {
	if !p.IsSetAvatarUrl() {
		return Creator_AvatarUrl_DEFAULT
	}
	return *p.AvatarUrl
}

var fieldIDToName_Creator = map[int16]string{
	1: "ID",
	2: "name",
	3: "avatarUrl",
}

func (p *Creator) IsSetID() bool {
	return p.ID != nil
}

func (p *Creator) IsSetName() bool {
	return p.Name != nil
}

func (p *Creator) IsSetAvatarUrl() bool {
	return p.AvatarUrl != nil
}

func (p *Creator) Read(iprot thrift.TProtocol) (err error) {
	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Creator[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Creator) ReadField1(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ID = _field
	return nil
}
func (p *Creator) ReadField2(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.Name = _field
	return nil
}
func (p *Creator) ReadField3(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.AvatarUrl = _field
	return nil
}

func (p *Creator) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Creator"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Creator) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetID() {
		if err = oprot.WriteFieldBegin("ID", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}
func (p *Creator) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetName() {
		if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Name); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}
func (p *Creator) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetAvatarUrl() {
		if err = oprot.WriteFieldBegin("avatarUrl", thrift.STRING, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.AvatarUrl); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Creator) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Creator(%+v)", *p)

}
