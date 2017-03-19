// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package sampling

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type SamplingStrategyType int64

const (
	SamplingStrategyType_PROBABILISTIC SamplingStrategyType = 0
	SamplingStrategyType_RATE_LIMITING SamplingStrategyType = 1
)

func (p SamplingStrategyType) String() string {
	switch p {
	case SamplingStrategyType_PROBABILISTIC:
		return "SamplingStrategyType_PROBABILISTIC"
	case SamplingStrategyType_RATE_LIMITING:
		return "SamplingStrategyType_RATE_LIMITING"
	}
	return "<UNSET>"
}

func SamplingStrategyTypeFromString(s string) (SamplingStrategyType, error) {
	switch s {
	case "SamplingStrategyType_PROBABILISTIC":
		return SamplingStrategyType_PROBABILISTIC, nil
	case "SamplingStrategyType_RATE_LIMITING":
		return SamplingStrategyType_RATE_LIMITING, nil
	}
	return SamplingStrategyType(0), fmt.Errorf("not a valid SamplingStrategyType string")
}

func SamplingStrategyTypePtr(v SamplingStrategyType) *SamplingStrategyType { return &v }

type ProbabilisticSamplingStrategy struct {
	SamplingRate float64 `thrift:"samplingRate,1,required" json:"samplingRate"`
}

func NewProbabilisticSamplingStrategy() *ProbabilisticSamplingStrategy {
	return &ProbabilisticSamplingStrategy{}
}

func (p *ProbabilisticSamplingStrategy) GetSamplingRate() float64 {
	return p.SamplingRate
}
func (p *ProbabilisticSamplingStrategy) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ProbabilisticSamplingStrategy) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.SamplingRate = v
	}
	return nil
}

func (p *ProbabilisticSamplingStrategy) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ProbabilisticSamplingStrategy"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ProbabilisticSamplingStrategy) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("samplingRate", thrift.DOUBLE, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:samplingRate: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.SamplingRate)); err != nil {
		return fmt.Errorf("%T.samplingRate (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:samplingRate: %s", p, err)
	}
	return err
}

func (p *ProbabilisticSamplingStrategy) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ProbabilisticSamplingStrategy(%+v)", *p)
}

type RateLimitingSamplingStrategy struct {
	MaxTracesPerSecond int16 `thrift:"maxTracesPerSecond,1,required" json:"maxTracesPerSecond"`
}

func NewRateLimitingSamplingStrategy() *RateLimitingSamplingStrategy {
	return &RateLimitingSamplingStrategy{}
}

func (p *RateLimitingSamplingStrategy) GetMaxTracesPerSecond() int16 {
	return p.MaxTracesPerSecond
}
func (p *RateLimitingSamplingStrategy) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *RateLimitingSamplingStrategy) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.MaxTracesPerSecond = v
	}
	return nil
}

func (p *RateLimitingSamplingStrategy) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RateLimitingSamplingStrategy"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *RateLimitingSamplingStrategy) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("maxTracesPerSecond", thrift.I16, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:maxTracesPerSecond: %s", p, err)
	}
	if err := oprot.WriteI16(int16(p.MaxTracesPerSecond)); err != nil {
		return fmt.Errorf("%T.maxTracesPerSecond (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:maxTracesPerSecond: %s", p, err)
	}
	return err
}

func (p *RateLimitingSamplingStrategy) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RateLimitingSamplingStrategy(%+v)", *p)
}

type OperationSamplingStrategy struct {
	Operation             string                         `thrift:"operation,1,required" json:"operation"`
	ProbabilisticSampling *ProbabilisticSamplingStrategy `thrift:"probabilisticSampling,2,required" json:"probabilisticSampling"`
}

func NewOperationSamplingStrategy() *OperationSamplingStrategy {
	return &OperationSamplingStrategy{}
}

func (p *OperationSamplingStrategy) GetOperation() string {
	return p.Operation
}

var OperationSamplingStrategy_ProbabilisticSampling_DEFAULT *ProbabilisticSamplingStrategy

func (p *OperationSamplingStrategy) GetProbabilisticSampling() *ProbabilisticSamplingStrategy {
	if !p.IsSetProbabilisticSampling() {
		return OperationSamplingStrategy_ProbabilisticSampling_DEFAULT
	}
	return p.ProbabilisticSampling
}
func (p *OperationSamplingStrategy) IsSetProbabilisticSampling() bool {
	return p.ProbabilisticSampling != nil
}

func (p *OperationSamplingStrategy) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *OperationSamplingStrategy) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Operation = v
	}
	return nil
}

func (p *OperationSamplingStrategy) ReadField2(iprot thrift.TProtocol) error {
	p.ProbabilisticSampling = &ProbabilisticSamplingStrategy{}
	if err := p.ProbabilisticSampling.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ProbabilisticSampling, err)
	}
	return nil
}

func (p *OperationSamplingStrategy) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("OperationSamplingStrategy"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *OperationSamplingStrategy) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("operation", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:operation: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Operation)); err != nil {
		return fmt.Errorf("%T.operation (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:operation: %s", p, err)
	}
	return err
}

func (p *OperationSamplingStrategy) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("probabilisticSampling", thrift.STRUCT, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:probabilisticSampling: %s", p, err)
	}
	if err := p.ProbabilisticSampling.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.ProbabilisticSampling, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:probabilisticSampling: %s", p, err)
	}
	return err
}

func (p *OperationSamplingStrategy) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OperationSamplingStrategy(%+v)", *p)
}

type PerOperationSamplingStrategies struct {
	DefaultSamplingProbability       float64                      `thrift:"defaultSamplingProbability,1,required" json:"defaultSamplingProbability"`
	DefaultLowerBoundTracesPerSecond float64                      `thrift:"defaultLowerBoundTracesPerSecond,2,required" json:"defaultLowerBoundTracesPerSecond"`
	PerOperationStrategies           []*OperationSamplingStrategy `thrift:"perOperationStrategies,3,required" json:"perOperationStrategies"`
}

func NewPerOperationSamplingStrategies() *PerOperationSamplingStrategies {
	return &PerOperationSamplingStrategies{}
}

func (p *PerOperationSamplingStrategies) GetDefaultSamplingProbability() float64 {
	return p.DefaultSamplingProbability
}

func (p *PerOperationSamplingStrategies) GetDefaultLowerBoundTracesPerSecond() float64 {
	return p.DefaultLowerBoundTracesPerSecond
}

func (p *PerOperationSamplingStrategies) GetPerOperationStrategies() []*OperationSamplingStrategy {
	return p.PerOperationStrategies
}
func (p *PerOperationSamplingStrategies) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PerOperationSamplingStrategies) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.DefaultSamplingProbability = v
	}
	return nil
}

func (p *PerOperationSamplingStrategies) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.DefaultLowerBoundTracesPerSecond = v
	}
	return nil
}

func (p *PerOperationSamplingStrategies) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*OperationSamplingStrategy, 0, size)
	p.PerOperationStrategies = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &OperationSamplingStrategy{}
		if err := _elem0.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem0, err)
		}
		p.PerOperationStrategies = append(p.PerOperationStrategies, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *PerOperationSamplingStrategies) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PerOperationSamplingStrategies"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PerOperationSamplingStrategies) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("defaultSamplingProbability", thrift.DOUBLE, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:defaultSamplingProbability: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.DefaultSamplingProbability)); err != nil {
		return fmt.Errorf("%T.defaultSamplingProbability (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:defaultSamplingProbability: %s", p, err)
	}
	return err
}

func (p *PerOperationSamplingStrategies) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("defaultLowerBoundTracesPerSecond", thrift.DOUBLE, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:defaultLowerBoundTracesPerSecond: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.DefaultLowerBoundTracesPerSecond)); err != nil {
		return fmt.Errorf("%T.defaultLowerBoundTracesPerSecond (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:defaultLowerBoundTracesPerSecond: %s", p, err)
	}
	return err
}

func (p *PerOperationSamplingStrategies) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("perOperationStrategies", thrift.LIST, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:perOperationStrategies: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.PerOperationStrategies)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.PerOperationStrategies {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:perOperationStrategies: %s", p, err)
	}
	return err
}

func (p *PerOperationSamplingStrategies) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PerOperationSamplingStrategies(%+v)", *p)
}

type SamplingStrategyResponse struct {
	StrategyType          SamplingStrategyType            `thrift:"strategyType,1,required" json:"strategyType"`
	ProbabilisticSampling *ProbabilisticSamplingStrategy  `thrift:"probabilisticSampling,2" json:"probabilisticSampling"`
	RateLimitingSampling  *RateLimitingSamplingStrategy   `thrift:"rateLimitingSampling,3" json:"rateLimitingSampling"`
	OperationSampling     *PerOperationSamplingStrategies `thrift:"operationSampling,4" json:"operationSampling"`
}

func NewSamplingStrategyResponse() *SamplingStrategyResponse {
	return &SamplingStrategyResponse{}
}

func (p *SamplingStrategyResponse) GetStrategyType() SamplingStrategyType {
	return p.StrategyType
}

var SamplingStrategyResponse_ProbabilisticSampling_DEFAULT *ProbabilisticSamplingStrategy

func (p *SamplingStrategyResponse) GetProbabilisticSampling() *ProbabilisticSamplingStrategy {
	if !p.IsSetProbabilisticSampling() {
		return SamplingStrategyResponse_ProbabilisticSampling_DEFAULT
	}
	return p.ProbabilisticSampling
}

var SamplingStrategyResponse_RateLimitingSampling_DEFAULT *RateLimitingSamplingStrategy

func (p *SamplingStrategyResponse) GetRateLimitingSampling() *RateLimitingSamplingStrategy {
	if !p.IsSetRateLimitingSampling() {
		return SamplingStrategyResponse_RateLimitingSampling_DEFAULT
	}
	return p.RateLimitingSampling
}

var SamplingStrategyResponse_OperationSampling_DEFAULT *PerOperationSamplingStrategies

func (p *SamplingStrategyResponse) GetOperationSampling() *PerOperationSamplingStrategies {
	if !p.IsSetOperationSampling() {
		return SamplingStrategyResponse_OperationSampling_DEFAULT
	}
	return p.OperationSampling
}
func (p *SamplingStrategyResponse) IsSetProbabilisticSampling() bool {
	return p.ProbabilisticSampling != nil
}

func (p *SamplingStrategyResponse) IsSetRateLimitingSampling() bool {
	return p.RateLimitingSampling != nil
}

func (p *SamplingStrategyResponse) IsSetOperationSampling() bool {
	return p.OperationSampling != nil
}

func (p *SamplingStrategyResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SamplingStrategyResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := SamplingStrategyType(v)
		p.StrategyType = temp
	}
	return nil
}

func (p *SamplingStrategyResponse) ReadField2(iprot thrift.TProtocol) error {
	p.ProbabilisticSampling = &ProbabilisticSamplingStrategy{}
	if err := p.ProbabilisticSampling.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ProbabilisticSampling, err)
	}
	return nil
}

func (p *SamplingStrategyResponse) ReadField3(iprot thrift.TProtocol) error {
	p.RateLimitingSampling = &RateLimitingSamplingStrategy{}
	if err := p.RateLimitingSampling.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.RateLimitingSampling, err)
	}
	return nil
}

func (p *SamplingStrategyResponse) ReadField4(iprot thrift.TProtocol) error {
	p.OperationSampling = &PerOperationSamplingStrategies{}
	if err := p.OperationSampling.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.OperationSampling, err)
	}
	return nil
}

func (p *SamplingStrategyResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SamplingStrategyResponse"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *SamplingStrategyResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("strategyType", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:strategyType: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.StrategyType)); err != nil {
		return fmt.Errorf("%T.strategyType (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:strategyType: %s", p, err)
	}
	return err
}

func (p *SamplingStrategyResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetProbabilisticSampling() {
		if err := oprot.WriteFieldBegin("probabilisticSampling", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:probabilisticSampling: %s", p, err)
		}
		if err := p.ProbabilisticSampling.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ProbabilisticSampling, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:probabilisticSampling: %s", p, err)
		}
	}
	return err
}

func (p *SamplingStrategyResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetRateLimitingSampling() {
		if err := oprot.WriteFieldBegin("rateLimitingSampling", thrift.STRUCT, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:rateLimitingSampling: %s", p, err)
		}
		if err := p.RateLimitingSampling.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.RateLimitingSampling, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:rateLimitingSampling: %s", p, err)
		}
	}
	return err
}

func (p *SamplingStrategyResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetOperationSampling() {
		if err := oprot.WriteFieldBegin("operationSampling", thrift.STRUCT, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:operationSampling: %s", p, err)
		}
		if err := p.OperationSampling.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.OperationSampling, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:operationSampling: %s", p, err)
		}
	}
	return err
}

func (p *SamplingStrategyResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SamplingStrategyResponse(%+v)", *p)
}