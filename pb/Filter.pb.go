// Code generated by protoc-gen-go.
// source: Filter.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FilterList_Operator int32

const (
	FilterList_MUST_PASS_ALL FilterList_Operator = 1
	FilterList_MUST_PASS_ONE FilterList_Operator = 2
)

var FilterList_Operator_name = map[int32]string{
	1: "MUST_PASS_ALL",
	2: "MUST_PASS_ONE",
}
var FilterList_Operator_value = map[string]int32{
	"MUST_PASS_ALL": 1,
	"MUST_PASS_ONE": 2,
}

func (x FilterList_Operator) Enum() *FilterList_Operator {
	p := new(FilterList_Operator)
	*p = x
	return p
}
func (x FilterList_Operator) String() string {
	return proto.EnumName(FilterList_Operator_name, int32(x))
}
func (x *FilterList_Operator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FilterList_Operator_value, data, "FilterList_Operator")
	if err != nil {
		return err
	}
	*x = FilterList_Operator(value)
	return nil
}
func (FilterList_Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{8, 0} }

type Filter struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SerializedFilter []byte  `protobuf:"bytes,2,opt,name=serialized_filter,json=serializedFilter" json:"serialized_filter,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Filter) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Filter) GetSerializedFilter() []byte {
	if m != nil {
		return m.SerializedFilter
	}
	return nil
}

type ColumnCountGetFilter struct {
	Limit            *int32 `protobuf:"varint,1,req,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ColumnCountGetFilter) Reset()                    { *m = ColumnCountGetFilter{} }
func (m *ColumnCountGetFilter) String() string            { return proto.CompactTextString(m) }
func (*ColumnCountGetFilter) ProtoMessage()               {}
func (*ColumnCountGetFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ColumnCountGetFilter) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type ColumnPaginationFilter struct {
	Limit            *int32 `protobuf:"varint,1,req,name=limit" json:"limit,omitempty"`
	Offset           *int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	ColumnOffset     []byte `protobuf:"bytes,3,opt,name=column_offset,json=columnOffset" json:"column_offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ColumnPaginationFilter) Reset()                    { *m = ColumnPaginationFilter{} }
func (m *ColumnPaginationFilter) String() string            { return proto.CompactTextString(m) }
func (*ColumnPaginationFilter) ProtoMessage()               {}
func (*ColumnPaginationFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ColumnPaginationFilter) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *ColumnPaginationFilter) GetOffset() int32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *ColumnPaginationFilter) GetColumnOffset() []byte {
	if m != nil {
		return m.ColumnOffset
	}
	return nil
}

type ColumnPrefixFilter struct {
	Prefix           []byte `protobuf:"bytes,1,req,name=prefix" json:"prefix,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ColumnPrefixFilter) Reset()                    { *m = ColumnPrefixFilter{} }
func (m *ColumnPrefixFilter) String() string            { return proto.CompactTextString(m) }
func (*ColumnPrefixFilter) ProtoMessage()               {}
func (*ColumnPrefixFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *ColumnPrefixFilter) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type ColumnRangeFilter struct {
	MinColumn          []byte `protobuf:"bytes,1,opt,name=min_column,json=minColumn" json:"min_column,omitempty"`
	MinColumnInclusive *bool  `protobuf:"varint,2,opt,name=min_column_inclusive,json=minColumnInclusive" json:"min_column_inclusive,omitempty"`
	MaxColumn          []byte `protobuf:"bytes,3,opt,name=max_column,json=maxColumn" json:"max_column,omitempty"`
	MaxColumnInclusive *bool  `protobuf:"varint,4,opt,name=max_column_inclusive,json=maxColumnInclusive" json:"max_column_inclusive,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *ColumnRangeFilter) Reset()                    { *m = ColumnRangeFilter{} }
func (m *ColumnRangeFilter) String() string            { return proto.CompactTextString(m) }
func (*ColumnRangeFilter) ProtoMessage()               {}
func (*ColumnRangeFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *ColumnRangeFilter) GetMinColumn() []byte {
	if m != nil {
		return m.MinColumn
	}
	return nil
}

func (m *ColumnRangeFilter) GetMinColumnInclusive() bool {
	if m != nil && m.MinColumnInclusive != nil {
		return *m.MinColumnInclusive
	}
	return false
}

func (m *ColumnRangeFilter) GetMaxColumn() []byte {
	if m != nil {
		return m.MaxColumn
	}
	return nil
}

func (m *ColumnRangeFilter) GetMaxColumnInclusive() bool {
	if m != nil && m.MaxColumnInclusive != nil {
		return *m.MaxColumnInclusive
	}
	return false
}

type CompareFilter struct {
	CompareOp        *CompareType `protobuf:"varint,1,req,name=compare_op,json=compareOp,enum=pb.CompareType" json:"compare_op,omitempty"`
	Comparator       *Comparator  `protobuf:"bytes,2,opt,name=comparator" json:"comparator,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CompareFilter) Reset()                    { *m = CompareFilter{} }
func (m *CompareFilter) String() string            { return proto.CompactTextString(m) }
func (*CompareFilter) ProtoMessage()               {}
func (*CompareFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *CompareFilter) GetCompareOp() CompareType {
	if m != nil && m.CompareOp != nil {
		return *m.CompareOp
	}
	return CompareType_LESS
}

func (m *CompareFilter) GetComparator() *Comparator {
	if m != nil {
		return m.Comparator
	}
	return nil
}

type DependentColumnFilter struct {
	CompareFilter       *CompareFilter `protobuf:"bytes,1,req,name=compare_filter,json=compareFilter" json:"compare_filter,omitempty"`
	ColumnFamily        []byte         `protobuf:"bytes,2,opt,name=column_family,json=columnFamily" json:"column_family,omitempty"`
	ColumnQualifier     []byte         `protobuf:"bytes,3,opt,name=column_qualifier,json=columnQualifier" json:"column_qualifier,omitempty"`
	DropDependentColumn *bool          `protobuf:"varint,4,opt,name=drop_dependent_column,json=dropDependentColumn" json:"drop_dependent_column,omitempty"`
	XXX_unrecognized    []byte         `json:"-"`
}

func (m *DependentColumnFilter) Reset()                    { *m = DependentColumnFilter{} }
func (m *DependentColumnFilter) String() string            { return proto.CompactTextString(m) }
func (*DependentColumnFilter) ProtoMessage()               {}
func (*DependentColumnFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *DependentColumnFilter) GetCompareFilter() *CompareFilter {
	if m != nil {
		return m.CompareFilter
	}
	return nil
}

func (m *DependentColumnFilter) GetColumnFamily() []byte {
	if m != nil {
		return m.ColumnFamily
	}
	return nil
}

func (m *DependentColumnFilter) GetColumnQualifier() []byte {
	if m != nil {
		return m.ColumnQualifier
	}
	return nil
}

func (m *DependentColumnFilter) GetDropDependentColumn() bool {
	if m != nil && m.DropDependentColumn != nil {
		return *m.DropDependentColumn
	}
	return false
}

type FamilyFilter struct {
	CompareFilter    *CompareFilter `protobuf:"bytes,1,req,name=compare_filter,json=compareFilter" json:"compare_filter,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FamilyFilter) Reset()                    { *m = FamilyFilter{} }
func (m *FamilyFilter) String() string            { return proto.CompactTextString(m) }
func (*FamilyFilter) ProtoMessage()               {}
func (*FamilyFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *FamilyFilter) GetCompareFilter() *CompareFilter {
	if m != nil {
		return m.CompareFilter
	}
	return nil
}

type FilterList struct {
	Operator         *FilterList_Operator `protobuf:"varint,1,req,name=operator,enum=pb.FilterList_Operator" json:"operator,omitempty"`
	Filters          []*Filter            `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *FilterList) Reset()                    { *m = FilterList{} }
func (m *FilterList) String() string            { return proto.CompactTextString(m) }
func (*FilterList) ProtoMessage()               {}
func (*FilterList) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *FilterList) GetOperator() FilterList_Operator {
	if m != nil && m.Operator != nil {
		return *m.Operator
	}
	return FilterList_MUST_PASS_ALL
}

func (m *FilterList) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type FilterWrapper struct {
	Filter           *Filter `protobuf:"bytes,1,req,name=filter" json:"filter,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FilterWrapper) Reset()                    { *m = FilterWrapper{} }
func (m *FilterWrapper) String() string            { return proto.CompactTextString(m) }
func (*FilterWrapper) ProtoMessage()               {}
func (*FilterWrapper) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *FilterWrapper) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type FirstKeyOnlyFilter struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FirstKeyOnlyFilter) Reset()                    { *m = FirstKeyOnlyFilter{} }
func (m *FirstKeyOnlyFilter) String() string            { return proto.CompactTextString(m) }
func (*FirstKeyOnlyFilter) ProtoMessage()               {}
func (*FirstKeyOnlyFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

type FirstKeyValueMatchingQualifiersFilter struct {
	Qualifiers       [][]byte `protobuf:"bytes,1,rep,name=qualifiers" json:"qualifiers,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FirstKeyValueMatchingQualifiersFilter) Reset()         { *m = FirstKeyValueMatchingQualifiersFilter{} }
func (m *FirstKeyValueMatchingQualifiersFilter) String() string { return proto.CompactTextString(m) }
func (*FirstKeyValueMatchingQualifiersFilter) ProtoMessage()    {}
func (*FirstKeyValueMatchingQualifiersFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{11}
}

func (m *FirstKeyValueMatchingQualifiersFilter) GetQualifiers() [][]byte {
	if m != nil {
		return m.Qualifiers
	}
	return nil
}

type FuzzyRowFilter struct {
	FuzzyKeysData    []*BytesBytesPair `protobuf:"bytes,1,rep,name=fuzzy_keys_data,json=fuzzyKeysData" json:"fuzzy_keys_data,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *FuzzyRowFilter) Reset()                    { *m = FuzzyRowFilter{} }
func (m *FuzzyRowFilter) String() string            { return proto.CompactTextString(m) }
func (*FuzzyRowFilter) ProtoMessage()               {}
func (*FuzzyRowFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *FuzzyRowFilter) GetFuzzyKeysData() []*BytesBytesPair {
	if m != nil {
		return m.FuzzyKeysData
	}
	return nil
}

type InclusiveStopFilter struct {
	StopRowKey       []byte `protobuf:"bytes,1,opt,name=stop_row_key,json=stopRowKey" json:"stop_row_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *InclusiveStopFilter) Reset()                    { *m = InclusiveStopFilter{} }
func (m *InclusiveStopFilter) String() string            { return proto.CompactTextString(m) }
func (*InclusiveStopFilter) ProtoMessage()               {}
func (*InclusiveStopFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{13} }

func (m *InclusiveStopFilter) GetStopRowKey() []byte {
	if m != nil {
		return m.StopRowKey
	}
	return nil
}

type KeyOnlyFilter struct {
	LenAsVal         *bool  `protobuf:"varint,1,req,name=len_as_val,json=lenAsVal" json:"len_as_val,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyOnlyFilter) Reset()                    { *m = KeyOnlyFilter{} }
func (m *KeyOnlyFilter) String() string            { return proto.CompactTextString(m) }
func (*KeyOnlyFilter) ProtoMessage()               {}
func (*KeyOnlyFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{14} }

func (m *KeyOnlyFilter) GetLenAsVal() bool {
	if m != nil && m.LenAsVal != nil {
		return *m.LenAsVal
	}
	return false
}

type MultipleColumnPrefixFilter struct {
	SortedPrefixes   [][]byte `protobuf:"bytes,1,rep,name=sorted_prefixes,json=sortedPrefixes" json:"sorted_prefixes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MultipleColumnPrefixFilter) Reset()                    { *m = MultipleColumnPrefixFilter{} }
func (m *MultipleColumnPrefixFilter) String() string            { return proto.CompactTextString(m) }
func (*MultipleColumnPrefixFilter) ProtoMessage()               {}
func (*MultipleColumnPrefixFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{15} }

func (m *MultipleColumnPrefixFilter) GetSortedPrefixes() [][]byte {
	if m != nil {
		return m.SortedPrefixes
	}
	return nil
}

type PageFilter struct {
	PageSize         *int64 `protobuf:"varint,1,req,name=page_size,json=pageSize" json:"page_size,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PageFilter) Reset()                    { *m = PageFilter{} }
func (m *PageFilter) String() string            { return proto.CompactTextString(m) }
func (*PageFilter) ProtoMessage()               {}
func (*PageFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{16} }

func (m *PageFilter) GetPageSize() int64 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return 0
}

type PrefixFilter struct {
	Prefix           []byte `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PrefixFilter) Reset()                    { *m = PrefixFilter{} }
func (m *PrefixFilter) String() string            { return proto.CompactTextString(m) }
func (*PrefixFilter) ProtoMessage()               {}
func (*PrefixFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{17} }

func (m *PrefixFilter) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type QualifierFilter struct {
	CompareFilter    *CompareFilter `protobuf:"bytes,1,req,name=compare_filter,json=compareFilter" json:"compare_filter,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *QualifierFilter) Reset()                    { *m = QualifierFilter{} }
func (m *QualifierFilter) String() string            { return proto.CompactTextString(m) }
func (*QualifierFilter) ProtoMessage()               {}
func (*QualifierFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{18} }

func (m *QualifierFilter) GetCompareFilter() *CompareFilter {
	if m != nil {
		return m.CompareFilter
	}
	return nil
}

type RandomRowFilter struct {
	Chance           *float32 `protobuf:"fixed32,1,req,name=chance" json:"chance,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RandomRowFilter) Reset()                    { *m = RandomRowFilter{} }
func (m *RandomRowFilter) String() string            { return proto.CompactTextString(m) }
func (*RandomRowFilter) ProtoMessage()               {}
func (*RandomRowFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{19} }

func (m *RandomRowFilter) GetChance() float32 {
	if m != nil && m.Chance != nil {
		return *m.Chance
	}
	return 0
}

type RowFilter struct {
	CompareFilter    *CompareFilter `protobuf:"bytes,1,req,name=compare_filter,json=compareFilter" json:"compare_filter,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RowFilter) Reset()                    { *m = RowFilter{} }
func (m *RowFilter) String() string            { return proto.CompactTextString(m) }
func (*RowFilter) ProtoMessage()               {}
func (*RowFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{20} }

func (m *RowFilter) GetCompareFilter() *CompareFilter {
	if m != nil {
		return m.CompareFilter
	}
	return nil
}

type SingleColumnValueExcludeFilter struct {
	SingleColumnValueFilter *SingleColumnValueFilter `protobuf:"bytes,1,req,name=single_column_value_filter,json=singleColumnValueFilter" json:"single_column_value_filter,omitempty"`
	XXX_unrecognized        []byte                   `json:"-"`
}

func (m *SingleColumnValueExcludeFilter) Reset()         { *m = SingleColumnValueExcludeFilter{} }
func (m *SingleColumnValueExcludeFilter) String() string { return proto.CompactTextString(m) }
func (*SingleColumnValueExcludeFilter) ProtoMessage()    {}
func (*SingleColumnValueExcludeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{21}
}

func (m *SingleColumnValueExcludeFilter) GetSingleColumnValueFilter() *SingleColumnValueFilter {
	if m != nil {
		return m.SingleColumnValueFilter
	}
	return nil
}

type SingleColumnValueFilter struct {
	ColumnFamily      []byte       `protobuf:"bytes,1,opt,name=column_family,json=columnFamily" json:"column_family,omitempty"`
	ColumnQualifier   []byte       `protobuf:"bytes,2,opt,name=column_qualifier,json=columnQualifier" json:"column_qualifier,omitempty"`
	CompareOp         *CompareType `protobuf:"varint,3,req,name=compare_op,json=compareOp,enum=pb.CompareType" json:"compare_op,omitempty"`
	Comparator        *Comparator  `protobuf:"bytes,4,req,name=comparator" json:"comparator,omitempty"`
	FilterIfMissing   *bool        `protobuf:"varint,5,opt,name=filter_if_missing,json=filterIfMissing" json:"filter_if_missing,omitempty"`
	LatestVersionOnly *bool        `protobuf:"varint,6,opt,name=latest_version_only,json=latestVersionOnly" json:"latest_version_only,omitempty"`
	XXX_unrecognized  []byte       `json:"-"`
}

func (m *SingleColumnValueFilter) Reset()                    { *m = SingleColumnValueFilter{} }
func (m *SingleColumnValueFilter) String() string            { return proto.CompactTextString(m) }
func (*SingleColumnValueFilter) ProtoMessage()               {}
func (*SingleColumnValueFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{22} }

func (m *SingleColumnValueFilter) GetColumnFamily() []byte {
	if m != nil {
		return m.ColumnFamily
	}
	return nil
}

func (m *SingleColumnValueFilter) GetColumnQualifier() []byte {
	if m != nil {
		return m.ColumnQualifier
	}
	return nil
}

func (m *SingleColumnValueFilter) GetCompareOp() CompareType {
	if m != nil && m.CompareOp != nil {
		return *m.CompareOp
	}
	return CompareType_LESS
}

func (m *SingleColumnValueFilter) GetComparator() *Comparator {
	if m != nil {
		return m.Comparator
	}
	return nil
}

func (m *SingleColumnValueFilter) GetFilterIfMissing() bool {
	if m != nil && m.FilterIfMissing != nil {
		return *m.FilterIfMissing
	}
	return false
}

func (m *SingleColumnValueFilter) GetLatestVersionOnly() bool {
	if m != nil && m.LatestVersionOnly != nil {
		return *m.LatestVersionOnly
	}
	return false
}

type SkipFilter struct {
	Filter           *Filter `protobuf:"bytes,1,req,name=filter" json:"filter,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SkipFilter) Reset()                    { *m = SkipFilter{} }
func (m *SkipFilter) String() string            { return proto.CompactTextString(m) }
func (*SkipFilter) ProtoMessage()               {}
func (*SkipFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{23} }

func (m *SkipFilter) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type TimestampsFilter struct {
	Timestamps       []int64 `protobuf:"varint,1,rep,packed,name=timestamps" json:"timestamps,omitempty"`
	CanHint          *bool   `protobuf:"varint,2,opt,name=can_hint,json=canHint" json:"can_hint,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TimestampsFilter) Reset()                    { *m = TimestampsFilter{} }
func (m *TimestampsFilter) String() string            { return proto.CompactTextString(m) }
func (*TimestampsFilter) ProtoMessage()               {}
func (*TimestampsFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{24} }

func (m *TimestampsFilter) GetTimestamps() []int64 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *TimestampsFilter) GetCanHint() bool {
	if m != nil && m.CanHint != nil {
		return *m.CanHint
	}
	return false
}

type ValueFilter struct {
	CompareFilter    *CompareFilter `protobuf:"bytes,1,req,name=compare_filter,json=compareFilter" json:"compare_filter,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ValueFilter) Reset()                    { *m = ValueFilter{} }
func (m *ValueFilter) String() string            { return proto.CompactTextString(m) }
func (*ValueFilter) ProtoMessage()               {}
func (*ValueFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{25} }

func (m *ValueFilter) GetCompareFilter() *CompareFilter {
	if m != nil {
		return m.CompareFilter
	}
	return nil
}

type WhileMatchFilter struct {
	Filter           *Filter `protobuf:"bytes,1,req,name=filter" json:"filter,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WhileMatchFilter) Reset()                    { *m = WhileMatchFilter{} }
func (m *WhileMatchFilter) String() string            { return proto.CompactTextString(m) }
func (*WhileMatchFilter) ProtoMessage()               {}
func (*WhileMatchFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{26} }

func (m *WhileMatchFilter) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type FilterAllFilter struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *FilterAllFilter) Reset()                    { *m = FilterAllFilter{} }
func (m *FilterAllFilter) String() string            { return proto.CompactTextString(m) }
func (*FilterAllFilter) ProtoMessage()               {}
func (*FilterAllFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{27} }

type RowRange struct {
	StartRow          []byte `protobuf:"bytes,1,opt,name=start_row,json=startRow" json:"start_row,omitempty"`
	StartRowInclusive *bool  `protobuf:"varint,2,opt,name=start_row_inclusive,json=startRowInclusive" json:"start_row_inclusive,omitempty"`
	StopRow           []byte `protobuf:"bytes,3,opt,name=stop_row,json=stopRow" json:"stop_row,omitempty"`
	StopRowInclusive  *bool  `protobuf:"varint,4,opt,name=stop_row_inclusive,json=stopRowInclusive" json:"stop_row_inclusive,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *RowRange) Reset()                    { *m = RowRange{} }
func (m *RowRange) String() string            { return proto.CompactTextString(m) }
func (*RowRange) ProtoMessage()               {}
func (*RowRange) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{28} }

func (m *RowRange) GetStartRow() []byte {
	if m != nil {
		return m.StartRow
	}
	return nil
}

func (m *RowRange) GetStartRowInclusive() bool {
	if m != nil && m.StartRowInclusive != nil {
		return *m.StartRowInclusive
	}
	return false
}

func (m *RowRange) GetStopRow() []byte {
	if m != nil {
		return m.StopRow
	}
	return nil
}

func (m *RowRange) GetStopRowInclusive() bool {
	if m != nil && m.StopRowInclusive != nil {
		return *m.StopRowInclusive
	}
	return false
}

type MultiRowRangeFilter struct {
	RowRangeList     []*RowRange `protobuf:"bytes,1,rep,name=row_range_list,json=rowRangeList" json:"row_range_list,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MultiRowRangeFilter) Reset()                    { *m = MultiRowRangeFilter{} }
func (m *MultiRowRangeFilter) String() string            { return proto.CompactTextString(m) }
func (*MultiRowRangeFilter) ProtoMessage()               {}
func (*MultiRowRangeFilter) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{29} }

func (m *MultiRowRangeFilter) GetRowRangeList() []*RowRange {
	if m != nil {
		return m.RowRangeList
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "pb.Filter")
	proto.RegisterType((*ColumnCountGetFilter)(nil), "pb.ColumnCountGetFilter")
	proto.RegisterType((*ColumnPaginationFilter)(nil), "pb.ColumnPaginationFilter")
	proto.RegisterType((*ColumnPrefixFilter)(nil), "pb.ColumnPrefixFilter")
	proto.RegisterType((*ColumnRangeFilter)(nil), "pb.ColumnRangeFilter")
	proto.RegisterType((*CompareFilter)(nil), "pb.CompareFilter")
	proto.RegisterType((*DependentColumnFilter)(nil), "pb.DependentColumnFilter")
	proto.RegisterType((*FamilyFilter)(nil), "pb.FamilyFilter")
	proto.RegisterType((*FilterList)(nil), "pb.FilterList")
	proto.RegisterType((*FilterWrapper)(nil), "pb.FilterWrapper")
	proto.RegisterType((*FirstKeyOnlyFilter)(nil), "pb.FirstKeyOnlyFilter")
	proto.RegisterType((*FirstKeyValueMatchingQualifiersFilter)(nil), "pb.FirstKeyValueMatchingQualifiersFilter")
	proto.RegisterType((*FuzzyRowFilter)(nil), "pb.FuzzyRowFilter")
	proto.RegisterType((*InclusiveStopFilter)(nil), "pb.InclusiveStopFilter")
	proto.RegisterType((*KeyOnlyFilter)(nil), "pb.KeyOnlyFilter")
	proto.RegisterType((*MultipleColumnPrefixFilter)(nil), "pb.MultipleColumnPrefixFilter")
	proto.RegisterType((*PageFilter)(nil), "pb.PageFilter")
	proto.RegisterType((*PrefixFilter)(nil), "pb.PrefixFilter")
	proto.RegisterType((*QualifierFilter)(nil), "pb.QualifierFilter")
	proto.RegisterType((*RandomRowFilter)(nil), "pb.RandomRowFilter")
	proto.RegisterType((*RowFilter)(nil), "pb.RowFilter")
	proto.RegisterType((*SingleColumnValueExcludeFilter)(nil), "pb.SingleColumnValueExcludeFilter")
	proto.RegisterType((*SingleColumnValueFilter)(nil), "pb.SingleColumnValueFilter")
	proto.RegisterType((*SkipFilter)(nil), "pb.SkipFilter")
	proto.RegisterType((*TimestampsFilter)(nil), "pb.TimestampsFilter")
	proto.RegisterType((*ValueFilter)(nil), "pb.ValueFilter")
	proto.RegisterType((*WhileMatchFilter)(nil), "pb.WhileMatchFilter")
	proto.RegisterType((*FilterAllFilter)(nil), "pb.FilterAllFilter")
	proto.RegisterType((*RowRange)(nil), "pb.RowRange")
	proto.RegisterType((*MultiRowRangeFilter)(nil), "pb.MultiRowRangeFilter")
	proto.RegisterEnum("pb.FilterList_Operator", FilterList_Operator_name, FilterList_Operator_value)
}

var fileDescriptor11 = []byte{
	// 1126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xd9, 0x52, 0x1b, 0x47,
	0x14, 0xad, 0x91, 0x58, 0xc4, 0x45, 0xeb, 0x80, 0x0d, 0xc1, 0x89, 0x8b, 0x9a, 0x6c, 0xc6, 0x49,
	0x54, 0x2e, 0xb9, 0x2a, 0x49, 0xe5, 0x0d, 0x61, 0x30, 0x94, 0x21, 0xe0, 0x11, 0xb1, 0xf3, 0x36,
	0xd5, 0x48, 0x2d, 0xd4, 0xc5, 0x68, 0x7a, 0x32, 0xdd, 0xc2, 0x88, 0x2f, 0xc8, 0x2f, 0xe4, 0xcd,
	0x2f, 0xf9, 0x84, 0x7c, 0x4a, 0xfe, 0xc7, 0xb7, 0xb7, 0x91, 0x64, 0xb0, 0xcb, 0xcb, 0x8b, 0x4a,
	0x7d, 0xee, 0xe9, 0xbb, 0x75, 0xdf, 0x33, 0x0d, 0xe5, 0x3d, 0x16, 0x4b, 0x9a, 0x35, 0xd3, 0x8c,
	0x4b, 0xee, 0x17, 0xd2, 0xb3, 0x8d, 0xe5, 0xfd, 0x36, 0x11, 0xd4, 0x00, 0x1b, 0xf5, 0x1d, 0x3e,
	0x4c, 0x49, 0x46, 0x24, 0xb7, 0x94, 0xe0, 0x00, 0x16, 0xcc, 0x16, 0xdf, 0x87, 0xb9, 0x84, 0x0c,
	0xe9, 0xba, 0xb7, 0x59, 0x78, 0xb0, 0x14, 0xea, 0xff, 0xfe, 0x0f, 0xd0, 0x10, 0x34, 0x63, 0x24,
	0x66, 0xd7, 0xb4, 0x17, 0xf5, 0x35, 0x71, 0xbd, 0xb0, 0xe9, 0x3d, 0x28, 0x87, 0xf5, 0x89, 0xc1,
	0x38, 0x08, 0x7e, 0x84, 0xd5, 0x1d, 0x1e, 0x8f, 0x86, 0xc9, 0x0e, 0x1f, 0x25, 0xf2, 0x29, 0x95,
	0xd6, 0xf1, 0x2a, 0xcc, 0xc7, 0x6c, 0xc8, 0xa4, 0xf6, 0x3c, 0x1f, 0x9a, 0x45, 0x70, 0x01, 0x77,
	0x0d, 0xfb, 0x84, 0x9c, 0xb3, 0x84, 0x48, 0xc6, 0x93, 0xf7, 0xf1, 0xfd, 0xbb, 0xb0, 0xc0, 0xfb,
	0x7d, 0x41, 0xa5, 0x8e, 0x3f, 0x1f, 0xda, 0x95, 0xff, 0x35, 0x54, 0xba, 0xda, 0x4f, 0x64, 0xcd,
	0x45, 0x9d, 0x5e, 0xd9, 0x80, 0xc7, 0x1a, 0xc3, 0xd4, 0x7c, 0x1b, 0x2c, 0xa3, 0x7d, 0x76, 0x65,
	0x03, 0xa1, 0xcb, 0x54, 0xaf, 0x75, 0xa4, 0x72, 0x68, 0x57, 0xc1, 0x7f, 0x1e, 0x34, 0x0c, 0x3d,
	0x24, 0xc9, 0x39, 0xb5, 0xec, 0xaf, 0x00, 0x86, 0x2c, 0x89, 0x8c, 0x5f, 0xdc, 0xa1, 0xa2, 0x2c,
	0x21, 0x62, 0x98, 0xfe, 0x23, 0x58, 0x9d, 0x98, 0x23, 0x96, 0x74, 0xe3, 0x91, 0x60, 0x97, 0x54,
	0x67, 0x5b, 0x0a, 0xfd, 0x9c, 0x78, 0xe0, 0x2c, 0xda, 0x21, 0xb9, 0x72, 0x0e, 0x8b, 0xd6, 0x21,
	0xb9, 0x9a, 0x72, 0x98, 0x9b, 0xa7, 0x1c, 0xce, 0x59, 0x87, 0x8e, 0x98, 0x3b, 0x0c, 0x38, 0x54,
	0xcc, 0xf9, 0xba, 0x94, 0x9b, 0x00, 0x5d, 0x03, 0x44, 0x3c, 0xd5, 0x45, 0x56, 0x5b, 0xb5, 0x66,
	0x7a, 0xd6, 0xb4, 0xb4, 0xd3, 0x71, 0x4a, 0xc3, 0x25, 0x4b, 0x39, 0x4e, 0x27, 0x7c, 0x75, 0x41,
	0x74, 0xe6, 0xcb, 0xad, 0xea, 0x84, 0xaf, 0xd0, 0x70, 0x8a, 0x11, 0xfc, 0xef, 0xc1, 0x9d, 0x27,
	0x34, 0xa5, 0x49, 0x8f, 0x26, 0xd2, 0x64, 0x63, 0x23, 0xff, 0x0a, 0x55, 0x17, 0xd9, 0xde, 0x1a,
	0x15, 0x7d, 0xb9, 0xd5, 0x98, 0x8a, 0x6e, 0xa8, 0x61, 0xa5, 0x3b, 0x93, 0xf3, 0xe4, 0x3c, 0xfb,
	0x64, 0xc8, 0xe2, 0xb1, 0xbd, 0x6e, 0xf6, 0x3c, 0xf7, 0x34, 0xe6, 0x6f, 0x41, 0xdd, 0x92, 0xfe,
	0x1a, 0xe1, 0x1d, 0xec, 0x33, 0x0c, 0x60, 0x1a, 0x58, 0x33, 0xf8, 0x73, 0x07, 0xfb, 0x2d, 0xb8,
	0xd3, 0xcb, 0x78, 0x1a, 0xf5, 0x5c, 0x9e, 0xae, 0xe1, 0xa6, 0x8f, 0x2b, 0xca, 0xf8, 0x56, 0x0d,
	0xc1, 0x3e, 0xce, 0x91, 0x0e, 0xf4, 0xb9, 0xd5, 0x04, 0xff, 0x78, 0x00, 0xe6, 0xef, 0x21, 0x13,
	0xd2, 0x7f, 0x0c, 0x25, 0x9e, 0x52, 0xd3, 0x5e, 0x73, 0x1c, 0x6b, 0xca, 0xc5, 0x84, 0xd1, 0x3c,
	0xb6, 0xe6, 0x30, 0x27, 0xfa, 0xdf, 0xc0, 0xa2, 0x89, 0x2a, 0xb0, 0x17, 0x45, 0x0c, 0x0b, 0x93,
	0x3d, 0xa1, 0x33, 0x05, 0x8f, 0xa0, 0xe4, 0xf6, 0xfa, 0x0d, 0xa8, 0x1c, 0xfd, 0xd1, 0x39, 0x8d,
	0x4e, 0xb6, 0x3b, 0x9d, 0x68, 0xfb, 0xf0, 0xb0, 0xee, 0xcd, 0x42, 0xc7, 0xbf, 0xef, 0xd6, 0x0b,
	0xc1, 0x63, 0xa8, 0x18, 0x27, 0x2f, 0x33, 0x92, 0xe2, 0x56, 0x3f, 0x80, 0x85, 0x99, 0xf2, 0xa6,
	0xe3, 0x58, 0x4b, 0xb0, 0x0a, 0xfe, 0x1e, 0xcb, 0x84, 0x7c, 0x46, 0xc7, 0xc7, 0x89, 0x6b, 0x50,
	0xf0, 0x14, 0xbe, 0x75, 0xe8, 0x0b, 0x12, 0x8f, 0xe8, 0x11, 0x91, 0xdd, 0x01, 0x4b, 0xce, 0xf3,
	0x63, 0x10, 0xb6, 0x93, 0xf7, 0x01, 0xf2, 0x13, 0x13, 0x18, 0xa6, 0x88, 0x47, 0x36, 0x85, 0x04,
	0x87, 0x50, 0xdd, 0x1b, 0x5d, 0x5f, 0x8f, 0x43, 0xfe, 0xca, 0xee, 0xf8, 0x0d, 0x6a, 0x7d, 0x85,
	0x44, 0x17, 0x74, 0x2c, 0xa2, 0x1e, 0x91, 0x44, 0x6f, 0x5b, 0x6e, 0xf9, 0x2a, 0xbb, 0xf6, 0x58,
	0x52, 0xa1, 0x7f, 0x4e, 0x08, 0xc3, 0xee, 0x6b, 0x2a, 0x66, 0x21, 0x9e, 0x20, 0x31, 0xf8, 0x05,
	0x56, 0xf2, 0xe9, 0xe8, 0x48, 0x9e, 0x5a, 0x97, 0x9b, 0x50, 0x16, 0xb8, 0x8a, 0x32, 0xfe, 0x4a,
	0x79, 0xb5, 0xb3, 0x0c, 0x0a, 0xc3, 0xb8, 0xb8, 0x3b, 0xf8, 0x09, 0x2a, 0x33, 0x05, 0xfa, 0x5f,
	0x02, 0xc4, 0x34, 0x89, 0x88, 0x88, 0x2e, 0x49, 0xac, 0xdb, 0x53, 0x0a, 0x4b, 0x88, 0x6c, 0x0b,
	0xac, 0x37, 0xd8, 0x85, 0x8d, 0xa3, 0x51, 0x2c, 0x59, 0x1a, 0xd3, 0x5b, 0x64, 0xe6, 0x7b, 0xa8,
	0x09, 0x9e, 0x49, 0x14, 0x50, 0xa3, 0x2f, 0xd4, 0x15, 0x5e, 0x35, 0xf0, 0x89, 0x45, 0x83, 0x2d,
	0x00, 0x14, 0x43, 0x37, 0x08, 0xf7, 0x60, 0x29, 0xc5, 0x55, 0x24, 0x50, 0x62, 0x75, 0xc4, 0x62,
	0x58, 0x52, 0x40, 0x07, 0xd7, 0xc1, 0x77, 0x50, 0x7e, 0xa7, 0x94, 0x79, 0x53, 0x52, 0xf6, 0x0c,
	0x6a, 0xf9, 0x19, 0x7c, 0xf6, 0x65, 0xde, 0x82, 0x1a, 0x0a, 0x62, 0x8f, 0x0f, 0x27, 0xa7, 0x83,
	0x71, 0xbb, 0x03, 0x92, 0x74, 0x4d, 0x86, 0x85, 0xd0, 0xae, 0xb0, 0x23, 0x4b, 0x13, 0xd2, 0xa7,
	0x47, 0xbc, 0x86, 0xfb, 0x1d, 0xbc, 0x46, 0xae, 0xad, 0xfa, 0x6e, 0xed, 0x5e, 0xe1, 0x89, 0xf6,
	0x5c, 0x97, 0xfe, 0x84, 0x0d, 0xa1, 0x19, 0x4e, 0x28, 0x2f, 0x15, 0x67, 0x36, 0xce, 0x3d, 0x15,
	0xe7, 0x86, 0x1f, 0x1b, 0x71, 0x4d, 0xdc, 0x6e, 0x08, 0xfe, 0x2d, 0xc0, 0xda, 0x3b, 0x36, 0xdd,
	0x14, 0x29, 0xef, 0x03, 0x45, 0xaa, 0x70, 0xbb, 0x48, 0xcd, 0x0a, 0x75, 0xf1, 0x23, 0x85, 0x7a,
	0x4e, 0x57, 0xf9, 0x1e, 0xa1, 0xf6, 0x1f, 0x42, 0xc3, 0x74, 0x24, 0x62, 0xfd, 0x68, 0xc8, 0x84,
	0xaa, 0x79, 0x7d, 0x5e, 0x0b, 0x60, 0xcd, 0x18, 0x0e, 0xfa, 0x47, 0x06, 0x46, 0xdf, 0x2b, 0x31,
	0xc1, 0x81, 0x92, 0xd1, 0x25, 0x4e, 0x24, 0x7e, 0x96, 0x23, 0x8e, 0x73, 0xb0, 0xbe, 0xa0, 0xd9,
	0x0d, 0x63, 0x7a, 0x61, 0x2c, 0x6a, 0x40, 0x50, 0x78, 0xa0, 0x73, 0xc1, 0xdc, 0x6c, 0x7d, 0x88,
	0x86, 0x3c, 0x87, 0xfa, 0x29, 0x1b, 0xa2, 0x1b, 0x32, 0x4c, 0x45, 0xbe, 0x0f, 0x64, 0x8e, 0xe9,
	0xf9, 0x28, 0xb6, 0x0b, 0x75, 0x2f, 0x9c, 0x42, 0xfd, 0x2f, 0xa0, 0xd4, 0x25, 0x49, 0x84, 0xc2,
	0x22, 0xed, 0x67, 0x75, 0x11, 0xd7, 0xfb, 0xb8, 0x44, 0x01, 0x5a, 0x9e, 0x3e, 0x9f, 0x4f, 0xbf,
	0x71, 0x3f, 0x43, 0xfd, 0xe5, 0x80, 0xc5, 0x46, 0xc1, 0x3e, 0xa2, 0xa6, 0x06, 0xd4, 0x0c, 0xb2,
	0x1d, 0xc7, 0xd6, 0xd5, 0x6b, 0x0f, 0x4a, 0x38, 0x04, 0xfa, 0x0d, 0xa1, 0xa6, 0x19, 0xab, 0xc8,
	0xa4, 0x12, 0x1d, 0x7b, 0x5b, 0x4a, 0x1a, 0x40, 0x86, 0x6a, 0x79, 0x6e, 0xbc, 0xf1, 0x74, 0x68,
	0x38, 0xda, 0xe4, 0xe5, 0x80, 0x8d, 0x70, 0x02, 0x66, 0x3f, 0x7b, 0x8b, 0x56, 0xbc, 0x7c, 0x7c,
	0xe9, 0xe4, 0xda, 0xf6, 0xf6, 0x9b, 0xa1, 0x6e, 0x49, 0x93, 0x17, 0xc3, 0x01, 0xac, 0x68, 0xe1,
	0x72, 0x69, 0xda, 0x82, 0x5b, 0x50, 0x55, 0xfb, 0x33, 0x05, 0x45, 0x31, 0x7e, 0x96, 0xac, 0xe4,
	0x96, 0x55, 0xe1, 0x8e, 0x1b, 0x96, 0x33, 0xfb, 0x4f, 0x7d, 0xb8, 0xda, 0x6d, 0x78, 0xc8, 0xb3,
	0xf3, 0x26, 0x49, 0x49, 0x77, 0x40, 0x9b, 0x03, 0xd2, 0xe3, 0x3c, 0x6d, 0x0e, 0xce, 0xf2, 0xb7,
	0xe7, 0xd9, 0xa8, 0xdf, 0x3c, 0xa7, 0x89, 0xfa, 0x42, 0xd1, 0x5e, 0xdb, 0xbe, 0x53, 0x4f, 0x94,
	0x45, 0xec, 0x7b, 0x7f, 0x7b, 0xde, 0x6b, 0xcf, 0x7b, 0x13, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd0,
	0xd7, 0x26, 0xbe, 0x0a, 0x00, 0x00,
}
