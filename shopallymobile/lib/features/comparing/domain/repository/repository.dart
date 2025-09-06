import 'package:dartz/dartz.dart';
import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';

abstract class Repository {
  Future<Either<Failure, ComparisonResultEntity>> compareProducts(List<ProductEntity> params);
  Future<Either<Failure, void>> removeProductFromCompare(String productId);
  Future<Either<Failure, List<ProductEntity>>> getProductsForComparison();
  Future<Either<Failure, void>> clearProducts();
}