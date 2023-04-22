
// Generated by Akka gRPC. DO NOT EDIT.
package service

import akka.annotation.ApiMayChange

import akka.grpc.AkkaGrpcGenerated


@AkkaGrpcGenerated
trait algo {
  
  
  def createAlgo(in: service.CreateAlgoReq): scala.concurrent.Future[service.AlgoReply]
  
  
  def queryAlgo(in: service.Empty): scala.concurrent.Future[service.AlgoReply]
  
  
  def dropAlgo(in: service.DropAlgoReq): scala.concurrent.Future[service.AlgoReply]
  
  
  def degree(in: service.BaseReq): scala.concurrent.Future[service.RankReply]
  
  
  def pagerank(in: service.BaseReq): scala.concurrent.Future[service.RankReply]
  
  
  def louvain(in: service.BaseReq): scala.concurrent.Future[service.ClusterReply]
  
  
  def shortestPath(in: service.ShortestPathReq): scala.concurrent.Future[service.ClusterReply]
  
  
  def avgShortestPath(in: service.BaseReq): scala.concurrent.Future[service.MetricsReply]
  
  
  def avgClustering(in: service.BaseReq): scala.concurrent.Future[service.MetricsReply]
  
  
  def custom(in: service.CustomAlgoReq): scala.concurrent.Future[service.CustomAlgoReply]
  
}



@AkkaGrpcGenerated
object algo extends akka.grpc.ServiceDescription {
  val name = "services.algo"

  val descriptor: com.google.protobuf.Descriptors.FileDescriptor =
    service.AlgoProto.javaDescriptor;

  object Serializers {
    import akka.grpc.scaladsl.ScalapbProtobufSerializer
    
    val CreateAlgoReqSerializer = new ScalapbProtobufSerializer(service.CreateAlgoReq.messageCompanion)
    
    val EmptySerializer = new ScalapbProtobufSerializer(service.Empty.messageCompanion)
    
    val DropAlgoReqSerializer = new ScalapbProtobufSerializer(service.DropAlgoReq.messageCompanion)
    
    val BaseReqSerializer = new ScalapbProtobufSerializer(service.BaseReq.messageCompanion)
    
    val ShortestPathReqSerializer = new ScalapbProtobufSerializer(service.ShortestPathReq.messageCompanion)
    
    val CustomAlgoReqSerializer = new ScalapbProtobufSerializer(service.CustomAlgoReq.messageCompanion)
    
    val AlgoReplySerializer = new ScalapbProtobufSerializer(service.AlgoReply.messageCompanion)
    
    val RankReplySerializer = new ScalapbProtobufSerializer(service.RankReply.messageCompanion)
    
    val ClusterReplySerializer = new ScalapbProtobufSerializer(service.ClusterReply.messageCompanion)
    
    val MetricsReplySerializer = new ScalapbProtobufSerializer(service.MetricsReply.messageCompanion)
    
    val CustomAlgoReplySerializer = new ScalapbProtobufSerializer(service.CustomAlgoReply.messageCompanion)
    
  }

  @ApiMayChange
  @AkkaGrpcGenerated
  object MethodDescriptors {
    import akka.grpc.internal.Marshaller
    import io.grpc.MethodDescriptor
    import Serializers._

    
    val createAlgoDescriptor: MethodDescriptor[service.CreateAlgoReq, service.AlgoReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "createAlgo"))
        .setRequestMarshaller(new Marshaller(CreateAlgoReqSerializer))
        .setResponseMarshaller(new Marshaller(AlgoReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val queryAlgoDescriptor: MethodDescriptor[service.Empty, service.AlgoReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "queryAlgo"))
        .setRequestMarshaller(new Marshaller(EmptySerializer))
        .setResponseMarshaller(new Marshaller(AlgoReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val dropAlgoDescriptor: MethodDescriptor[service.DropAlgoReq, service.AlgoReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "dropAlgo"))
        .setRequestMarshaller(new Marshaller(DropAlgoReqSerializer))
        .setResponseMarshaller(new Marshaller(AlgoReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val degreeDescriptor: MethodDescriptor[service.BaseReq, service.RankReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "degree"))
        .setRequestMarshaller(new Marshaller(BaseReqSerializer))
        .setResponseMarshaller(new Marshaller(RankReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val pagerankDescriptor: MethodDescriptor[service.BaseReq, service.RankReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "pagerank"))
        .setRequestMarshaller(new Marshaller(BaseReqSerializer))
        .setResponseMarshaller(new Marshaller(RankReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val louvainDescriptor: MethodDescriptor[service.BaseReq, service.ClusterReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "louvain"))
        .setRequestMarshaller(new Marshaller(BaseReqSerializer))
        .setResponseMarshaller(new Marshaller(ClusterReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val shortestPathDescriptor: MethodDescriptor[service.ShortestPathReq, service.ClusterReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "shortestPath"))
        .setRequestMarshaller(new Marshaller(ShortestPathReqSerializer))
        .setResponseMarshaller(new Marshaller(ClusterReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val avgShortestPathDescriptor: MethodDescriptor[service.BaseReq, service.MetricsReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "avgShortestPath"))
        .setRequestMarshaller(new Marshaller(BaseReqSerializer))
        .setResponseMarshaller(new Marshaller(MetricsReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val avgClusteringDescriptor: MethodDescriptor[service.BaseReq, service.MetricsReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "avgClustering"))
        .setRequestMarshaller(new Marshaller(BaseReqSerializer))
        .setResponseMarshaller(new Marshaller(MetricsReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
    val customDescriptor: MethodDescriptor[service.CustomAlgoReq, service.CustomAlgoReply] =
      MethodDescriptor.newBuilder()
        .setType(
   MethodDescriptor.MethodType.UNARY 
  
  
  
)
        .setFullMethodName(MethodDescriptor.generateFullMethodName("services.algo", "custom"))
        .setRequestMarshaller(new Marshaller(CustomAlgoReqSerializer))
        .setResponseMarshaller(new Marshaller(CustomAlgoReplySerializer))
        .setSampledToLocalTracing(true)
        .build()
    
  }
}
