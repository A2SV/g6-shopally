part of 'saved_product_bloc.dart';

sealed class SavedProductEvent extends Equatable {
  const SavedProductEvent();

  @override
  List<Object> get props => [];
}


final class SaveProductEvent extends SavedProductEvent {
  final Product product;

  const SaveProductEvent(this.product);

  @override
  List<Object> get props => [product];
}

final class RemoveProductEvent extends SavedProductEvent {
  final String productId;

  const RemoveProductEvent(this.productId);

  @override
  List<Object> get props => [productId];
}

final class LoadSavedProductsEvent extends SavedProductEvent {
  const LoadSavedProductsEvent();
}


//////////   add to compare event /////////////
final class AddToCompareEvent extends SavedProductEvent {
  final Product product;

  const AddToCompareEvent(this.product);

  @override
  List<Object> get props => [product];
}

//////////   remove from compare event /////////////
final class RemoveFromCompareEvent extends SavedProductEvent {
  final String productId;

  const RemoveFromCompareEvent(this.productId);

  @override
  List<Object> get props => [productId];
}