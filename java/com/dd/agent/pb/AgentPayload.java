// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/logs/agent_logs_payload.proto

package com.dd.agent.pb;

public final class AgentPayload {
  private AgentPayload() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_pb_Log_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_pb_Log_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n#proto/logs/agent_logs_payload.proto\022\002p" +
      "b\"z\n\003Log\022\017\n\007message\030\001 \001(\t\022\016\n\006status\030\002 \001(" +
      "\t\022\021\n\ttimestamp\030\003 \001(\003\022\020\n\010hostname\030\004 \001(\t\022\017" +
      "\n\007service\030\005 \001(\t\022\016\n\006source\030\006 \001(\t\022\014\n\004tags\030" +
      "\007 \003(\tBI\n\017com.dd.agent.pbB\014AgentPayloadP\001" +
      "Z&github.com/DataDog/agent-payload/v5/pb" +
      "b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_pb_Log_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_pb_Log_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_pb_Log_descriptor,
        new java.lang.String[] { "Message", "Status", "Timestamp", "Hostname", "Service", "Source", "Tags", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
