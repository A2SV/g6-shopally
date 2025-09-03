part of 'chat_bloc.dart';

@immutable
sealed class ChatState {
  const ChatState();
}

final class ChatInitial extends ChatState {
  const ChatInitial();
}

class ResponseLoadingState extends ChatState {
  const ResponseLoadingState();
}

class ProductsLoadedState extends ChatState {
  final List<ProductEntity> products;
  const ProductsLoadedState({this.products = const []});
}
class ChatErrorState extends ChatState {
  final String message;
  const ChatErrorState(this.message);
}
class ProductDetailState extends ChatState {
  final ProductEntity product;
  const ProductDetailState(this.product);
}
class ProductDetailLoadingState extends ChatState {
  const ProductDetailLoadingState();
}
class ProductDetailErrorState extends ChatState {
  final String message;
  const ProductDetailErrorState(this.message);
}
class ChatLoadingState extends ChatState {
  const ChatLoadingState();
}
class ResponseErrorState extends ChatState {
  final String message;
  const ResponseErrorState(this.message);
}

class LoadingState extends ChatState {
  const LoadingState();
}
class ResponseLoadedState extends ChatState {
  final List<ProductEntity> products;
  const ResponseLoadedState({this.products = const []});
}