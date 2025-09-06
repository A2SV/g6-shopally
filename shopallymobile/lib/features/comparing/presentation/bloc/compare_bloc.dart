import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
import 'package:shopallymobile/features/comparing/core/usecases/usecases.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/clear_products_useCase.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/compare_products_usecase.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/get_products_for_comparison_useCase.dart';
import 'package:shopallymobile/features/comparing/domain/usecases/remove_from_compare_usecase.dart';

part 'compare_event.dart';
part 'compare_state.dart';

class CompareBloc extends Bloc<CompareEvent, CompareState> {
  final CompareProductsUseCase compareProductsUseCase;
  final GetProductsForComparison getProductsForComparisonUseCase;
  final RemoveFromCompareUseCase removeProductFromComparisonUseCase;
  final ClearProductsUseCase clearProductsUseCase;
  CompareBloc({
    required this.compareProductsUseCase,
    required this.getProductsForComparisonUseCase,
    required this.removeProductFromComparisonUseCase,
    required this.clearProductsUseCase,
  }) : super(CompareInitial()) {
    on<FetchProductForComparisonEvent>((event, emit) async {
      emit(CompareLoading());
      final result = await getProductsForComparisonUseCase(NoParams());

      result.fold(
        (failure) {
          emit(CompareError(failure.message));
        },
        (products) {
          if (products.isEmpty) {
            emit(ComparisonEmpty());
          } else {
            emit(ComparingProductLoaded(products));
          }
        },
      );
    });

    on<RemoveProductFromComparisonEvent>((event, emit) async {
      emit(CompareLoading());
      final result = await removeProductFromComparisonUseCase(event.productId);
      await result.fold(
        (failure) async {
          emit(CompareError(failure.message));
        },
        (_) async {
          add(FetchProductForComparisonEvent());
        },
      );
    });

    on<CompareProductsEvent>((event, emit) async {
      emit(CompareLoading());
      final result = await compareProductsUseCase(event.products);
      result.fold(
        (failure) {
          emit(CompareError(failure.message));
        },
        (comparison) {
          emit(ComparisonResult(comparison));
        },
      );
    });

    on<ClearProductsFromComparisonEvent>((event, emit) async {
      emit(CompareLoading());
      final result = await clearProductsUseCase(NoParams());
      await result.fold(
        (failure) async {
          emit(CompareError(failure.message));
        },
        (_) async {
          add(FetchProductForComparisonEvent());
        },
      );
    });
  }
}
