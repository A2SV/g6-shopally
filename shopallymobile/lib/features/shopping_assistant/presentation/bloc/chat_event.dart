part of 'chat_bloc.dart';

@immutable
sealed class ChatEvent {}

class SendPromptRequested extends ChatEvent {
  final String prompt;

  SendPromptRequested(this.prompt);
}

class ProductDetailRequested extends ChatEvent {
  final String productId;
  
  ProductDetailRequested(this.productId);
}