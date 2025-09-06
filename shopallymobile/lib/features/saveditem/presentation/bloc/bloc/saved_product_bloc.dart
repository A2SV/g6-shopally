import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';

part 'saved_product_event.dart';
part 'saved_product_state.dart';

class SavedProductBloc extends Bloc<SavedProductEvent, SavedProductState> {
  final SavedItemsRepository savedItemsRepository;

  SavedProductBloc(this.savedItemsRepository)
    : super(const SavedProductInitial()) {
    on<SaveProductEvent>(_onSaveProduct);
    on<RemoveProductEvent>(_onRemoveProduct);
    on<LoadSavedProductsEvent>(_onLoadSavedProducts);
    on<AddToCompareEvent>(_onAddToCompare);
    on<RemoveFromCompareEvent>(_onRemoveFromCompare);
  }

  Future<void> _onSaveProduct(
    SaveProductEvent event,
    Emitter<SavedProductState> emit,
  ) async {
    emit(const SavedProductLoading());
    final result = await savedItemsRepository.saveProduct(event.product);
    await result.fold(
      (failure) async => emit(SavedProductError(failure.message)),
      (_) async {
        // After successful save, reload the saved products list
        final loadResult = await savedItemsRepository.getSavedItems();
        await loadResult.fold(
          (failure) async => emit(SavedProductError(failure.message)),
          (products) async => emit(SavedProductLoaded(products)),
        );
      },
    );
  }

  Future<void> _onRemoveProduct(
    RemoveProductEvent event,
    Emitter<SavedProductState> emit,
  ) async {
    emit(const SavedProductLoading());
    final result = await savedItemsRepository.removeProduct(event.productId);
    await result.fold(
      (failure) async => emit(RemoveProductError(failure.message)),
      (_) async {
        // After successful removal, reload the saved products list
        final loadResult = await savedItemsRepository.getSavedItems();
        await loadResult.fold(
          (failure) async => emit(SavedProductError(failure.message)),
          (products) async => emit(SavedProductLoaded(products)),
        );
      },
    );
  }

  Future<void> _onLoadSavedProducts(
    LoadSavedProductsEvent event,
    Emitter<SavedProductState> emit,
  ) async {
    emit(const SavedProductLoading());
    try {
      final result = await savedItemsRepository.getSavedItems();

      result.fold(
        (failure) => emit(SavedProductError(failure.message)),
        (products) => emit(SavedProductLoaded(products)),
      );
    } catch (e) {
      emit(SavedProductError('An unexpected error occurred: $e'));
    }
  }

  Future<void> _onAddToCompare(
    AddToCompareEvent event,
    Emitter<SavedProductState> emit,
  ) async {
    emit(const SavedProductLoading());
    final result = await savedItemsRepository.addtoCompare(event.product);
    await result.fold(
      (failure) async => emit(AddToCompareError(failure.message)),
      (_) async {
        // After successful addition to compare, reload the saved products list
        final loadResult = await savedItemsRepository.getSavedItems();
        await loadResult.fold(
          (failure) async => emit(SavedProductError(failure.message)),
          (products) async => emit(SavedProductLoaded(products)),
        );
      },
    );
  }

  Future<void> _onRemoveFromCompare(
    RemoveFromCompareEvent event,
    Emitter<SavedProductState> emit,
  ) async {
    emit(const SavedProductLoading());
    final result = await savedItemsRepository.removefromCompare(
      event.productId,
    );
    result.fold(
      (failure) => emit(RemoveFromCompareError(failure.message)),
      (_) => emit(const RemoveFromCompareSuccess()),
    );
  }
}
