// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: movie.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImdbId string `protobuf:"bytes,1,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *DetailRequest) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Search string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Data) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Data) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Data) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Data) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*Data     `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize    int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	CurrentPage int32 `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	TotalPage   int32 `protobuf:"varint,3,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	TotalResult int32 `protobuf:"varint,4,opt,name=total_result,json=totalResult,proto3" json:"total_result,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{4}
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *Pagination) GetTotalResult() int32 {
	if x != nil {
		return x.TotalResult
	}
	return 0
}

type Ratings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Ratings) Reset() {
	*x = Ratings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ratings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ratings) ProtoMessage() {}

func (x *Ratings) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ratings.ProtoReflect.Descriptor instead.
func (*Ratings) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{5}
}

func (x *Ratings) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Ratings) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DetailMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string     `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year       string     `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	Rated      string     `protobuf:"bytes,3,opt,name=Rated,proto3" json:"Rated,omitempty"`
	Released   string     `protobuf:"bytes,4,opt,name=Released,proto3" json:"Released,omitempty"`
	Runtime    string     `protobuf:"bytes,5,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Genre      string     `protobuf:"bytes,6,opt,name=Genre,proto3" json:"Genre,omitempty"`
	Director   string     `protobuf:"bytes,7,opt,name=Director,proto3" json:"Director,omitempty"`
	Writer     string     `protobuf:"bytes,8,opt,name=Writer,proto3" json:"Writer,omitempty"`
	Actors     string     `protobuf:"bytes,9,opt,name=Actors,proto3" json:"Actors,omitempty"`
	Plot       string     `protobuf:"bytes,10,opt,name=Plot,proto3" json:"Plot,omitempty"`
	Language   string     `protobuf:"bytes,11,opt,name=Language,proto3" json:"Language,omitempty"`
	Country    string     `protobuf:"bytes,12,opt,name=Country,proto3" json:"Country,omitempty"`
	Awards     string     `protobuf:"bytes,13,opt,name=Awards,proto3" json:"Awards,omitempty"`
	Poster     string     `protobuf:"bytes,14,opt,name=Poster,proto3" json:"Poster,omitempty"`
	Ratings    []*Ratings `protobuf:"bytes,15,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Metascore  string     `protobuf:"bytes,16,opt,name=Metascore,proto3" json:"Metascore,omitempty"`
	ImdbRating string     `protobuf:"bytes,17,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string     `protobuf:"bytes,18,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string     `protobuf:"bytes,19,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string     `protobuf:"bytes,20,opt,name=Type,proto3" json:"Type,omitempty"`
	DVD        string     `protobuf:"bytes,21,opt,name=DVD,proto3" json:"DVD,omitempty"`
	BoxOffice  string     `protobuf:"bytes,22,opt,name=BoxOffice,proto3" json:"BoxOffice,omitempty"`
	Production string     `protobuf:"bytes,23,opt,name=Production,proto3" json:"Production,omitempty"`
	Website    string     `protobuf:"bytes,24,opt,name=Website,proto3" json:"Website,omitempty"`
	Response   string     `protobuf:"bytes,25,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *DetailMovie) Reset() {
	*x = DetailMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailMovie) ProtoMessage() {}

func (x *DetailMovie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailMovie.ProtoReflect.Descriptor instead.
func (*DetailMovie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{6}
}

func (x *DetailMovie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DetailMovie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *DetailMovie) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *DetailMovie) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *DetailMovie) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *DetailMovie) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *DetailMovie) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *DetailMovie) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *DetailMovie) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *DetailMovie) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *DetailMovie) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *DetailMovie) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *DetailMovie) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *DetailMovie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *DetailMovie) GetRatings() []*Ratings {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *DetailMovie) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *DetailMovie) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *DetailMovie) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *DetailMovie) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *DetailMovie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DetailMovie) GetDVD() string {
	if x != nil {
		return x.DVD
	}
	return ""
}

func (x *DetailMovie) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *DetailMovie) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *DetailMovie) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *DetailMovie) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x22, 0x39,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x74, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x62, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x37, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x97, 0x05,
	0x0a, 0x0b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x41, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x64,
	0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x64, 0x62,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x56, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x44, 0x56, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6c, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_movie_proto_goTypes = []interface{}{
	(*DetailRequest)(nil), // 0: Movie.DetailRequest
	(*ListRequest)(nil),   // 1: Movie.ListRequest
	(*Data)(nil),          // 2: Movie.data
	(*ListResponse)(nil),  // 3: Movie.ListResponse
	(*Pagination)(nil),    // 4: Movie.Pagination
	(*Ratings)(nil),       // 5: Movie.Ratings
	(*DetailMovie)(nil),   // 6: Movie.DetailMovie
}
var file_movie_proto_depIdxs = []int32{
	2, // 0: Movie.ListResponse.data:type_name -> Movie.data
	4, // 1: Movie.ListResponse.pagination:type_name -> Movie.Pagination
	5, // 2: Movie.DetailMovie.Ratings:type_name -> Movie.Ratings
	1, // 3: Movie.Movie.List:input_type -> Movie.ListRequest
	0, // 4: Movie.Movie.Detail:input_type -> Movie.DetailRequest
	3, // 5: Movie.Movie.List:output_type -> Movie.ListResponse
	6, // 6: Movie.Movie.Detail:output_type -> Movie.DetailMovie
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ratings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailMovie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieClient is the client API for Movie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailMovie, error)
}

type movieClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieClient(cc grpc.ClientConnInterface) MovieClient {
	return &movieClient{cc}
}

func (c *movieClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/Movie.Movie/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailMovie, error) {
	out := new(DetailMovie)
	err := c.cc.Invoke(ctx, "/Movie.Movie/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServer is the server API for Movie service.
type MovieServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailMovie, error)
}

// UnimplementedMovieServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServer struct {
}

func (*UnimplementedMovieServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedMovieServer) Detail(context.Context, *DetailRequest) (*DetailMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}

func RegisterMovieServer(s *grpc.Server, srv MovieServer) {
	s.RegisterService(&_Movie_serviceDesc, srv)
}

func _Movie_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Movie.Movie/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Movie.Movie/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Movie_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Movie.Movie",
	HandlerType: (*MovieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Movie_List_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Movie_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
