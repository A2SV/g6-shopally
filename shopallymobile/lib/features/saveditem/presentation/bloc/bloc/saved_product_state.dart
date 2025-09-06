part of 'saved_product_bloc.dart';

sealed class SavedProductState extends Equatable {
  const SavedProductState();
  
  @override
  List<Object> get props => [];
}

final class SavedProductInitial extends SavedProductState {
  const SavedProductInitial();
}
final class SavedProductLoading extends SavedProductState {
  const SavedProductLoading();
}

final class SaveProductSuccess extends SavedProductState {
  const SaveProductSuccess();
}

final class SavedProductLoaded extends SavedProductState {
  final List<Product> products;

  const SavedProductLoaded(this.products);

  @override
  List<Object> get props => [products];
}


final class SaveProductError extends SavedProductState {
  final String message;

  const SaveProductError(this.message);

  @override
  List<Object> get props => [message];
}
////////////////////////////Remove////////////////////////////////

final class RemoveProductError extends SavedProductState {
  final String message;

  const RemoveProductError(this.message);

  @override
  List<Object> get props => [message];
}
final class RemoveProductSuccess extends SavedProductState {
  const RemoveProductSuccess();
}

////////////////////////add to compare ///////
final class AddToCompareError extends SavedProductState {
  final String message;

  const AddToCompareError(this.message);

  @override
  List<Object> get props => [message];
}
final class AddToCompareSuccess extends SavedProductState {
  const AddToCompareSuccess();
}

////////////////////////remove from compare ///////
final class RemoveFromCompareError extends SavedProductState {
  final String message;

  const RemoveFromCompareError(this.message);

  @override
  List<Object> get props => [message];
}
final class RemoveFromCompareSuccess extends SavedProductState {
  const RemoveFromCompareSuccess();
}
///
// Example state classesR
class SavedProductError extends SavedProductState {
  final String message;
  SavedProductError(this.message);
}