part of 'compare_bloc.dart';

@immutable
sealed class CompareEvent {}

class FetchProductForComparisonEvent extends CompareEvent {}

class RemoveProductFromComparisonEvent extends CompareEvent {
  final String productId;
  RemoveProductFromComparisonEvent({ required this.productId});
}

class CompareProductsEvent extends CompareEvent {
  final List<ProductEntity> products;
  CompareProductsEvent(this.products);
}

class ClearProductsFromComparisonEvent extends CompareEvent{}