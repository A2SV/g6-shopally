import 'package:dartz/dartz.dart';
import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/core/usecases/usecases.dart' show UseCase;
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';


class CompareProductsUseCase
    extends UseCase<ComparisonResultEntity, List<ProductEntity>> {
  final Repository repository;
  CompareProductsUseCase({required this.repository});
  @override
  Future<Either<Failure, ComparisonResultEntity>> call(
    List<ProductEntity> params,
  ) {
    return repository.compareProducts(params);
  }
}
