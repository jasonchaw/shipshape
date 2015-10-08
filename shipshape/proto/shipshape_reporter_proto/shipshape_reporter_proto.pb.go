// Code generated by protoc-gen-go.
// source: shipshape/proto/shipshape_reporter.proto
// DO NOT EDIT!

/*
Package shipshape_reporter_proto_go_src is a generated protocol buffer package.

It is generated from these files:
	shipshape/proto/shipshape_reporter.proto

It has these top-level messages:
	ReportNotesRequest
	ReportNotesResponse
	ReportAnalyzerStatusRequest
	ReportAnalyzerStatusResponse
*/
package shipshape_reporter_proto_go_src

import proto "github.com/golang/protobuf/proto"
import math "math"
import shipshape_proto1 "github.com/google/shipshape/shipshape/proto/note_proto"

// discarding unused import shipshape_proto2 "github.com/google/shipshape/shipshape/proto/shipshape_context_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AnalyzerStatus int32

const (
	AnalyzerStatus_RUNNING   AnalyzerStatus = 1
	AnalyzerStatus_COMPLETED AnalyzerStatus = 2
	AnalyzerStatus_FAILED    AnalyzerStatus = 3
)

var AnalyzerStatus_name = map[int32]string{
	1: "RUNNING",
	2: "COMPLETED",
	3: "FAILED",
}
var AnalyzerStatus_value = map[string]int32{
	"RUNNING":   1,
	"COMPLETED": 2,
	"FAILED":    3,
}

func (x AnalyzerStatus) Enum() *AnalyzerStatus {
	p := new(AnalyzerStatus)
	*p = x
	return p
}
func (x AnalyzerStatus) String() string {
	return proto.EnumName(AnalyzerStatus_name, int32(x))
}
func (x *AnalyzerStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AnalyzerStatus_value, data, "AnalyzerStatus")
	if err != nil {
		return err
	}
	*x = AnalyzerStatus(value)
	return nil
}

type ReportNotesRequest struct {
	// All notes should have the same category.
	Notes            []*shipshape_proto1.Note `protobuf:"bytes,1,rep,name=notes" json:"notes,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ReportNotesRequest) Reset()         { *m = ReportNotesRequest{} }
func (m *ReportNotesRequest) String() string { return proto.CompactTextString(m) }
func (*ReportNotesRequest) ProtoMessage()    {}

func (m *ReportNotesRequest) GetNotes() []*shipshape_proto1.Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

type ReportNotesResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReportNotesResponse) Reset()         { *m = ReportNotesResponse{} }
func (m *ReportNotesResponse) String() string { return proto.CompactTextString(m) }
func (*ReportNotesResponse) ProtoMessage()    {}

type ReportAnalyzerStatusRequest struct {
	Category         *string         `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
	Status           *AnalyzerStatus `protobuf:"varint,2,opt,name=status,enum=shipshape_proto.AnalyzerStatus" json:"status,omitempty"`
	Message          *string         `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ReportAnalyzerStatusRequest) Reset()         { *m = ReportAnalyzerStatusRequest{} }
func (m *ReportAnalyzerStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ReportAnalyzerStatusRequest) ProtoMessage()    {}

func (m *ReportAnalyzerStatusRequest) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *ReportAnalyzerStatusRequest) GetStatus() AnalyzerStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return AnalyzerStatus_RUNNING
}

func (m *ReportAnalyzerStatusRequest) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ReportAnalyzerStatusResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReportAnalyzerStatusResponse) Reset()         { *m = ReportAnalyzerStatusResponse{} }
func (m *ReportAnalyzerStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ReportAnalyzerStatusResponse) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("shipshape_proto.AnalyzerStatus", AnalyzerStatus_name, AnalyzerStatus_value)
}