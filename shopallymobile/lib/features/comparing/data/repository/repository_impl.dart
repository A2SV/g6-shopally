import 'package:dartz/dartz.dart';
import 'package:shopallymobile/features/comparing/core/network/network_info.dart';
import 'package:shopallymobile/features/comparing/core/errors/exception.dart';
import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/data/datasources/local_data_source.dart';
import 'package:shopallymobile/features/comparing/data/datasources/remote_data_source.dart';
import 'package:shopallymobile/features/comparing/data/model/product_model.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';


import '../model/comparison_result_model.dart';

class RepositoryImpl extends Repository {
  final LocalDataSource localDataSource;
  final RemoteDataSource remoteDataSource;
  final NetworkInfo networkInfo;

  RepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
    required this.networkInfo,
  });
  @override
  Future<Either<Failure, ComparisonResultEntity>> compareProducts(
    List<ProductEntity> products,
  ) async {
    if (await networkInfo.isConnected) {
      try {
        final List<ProductModel> productModels = products
            .map((product) => ProductModel.fromEntity(product))
            .toList();
        final ComparisonResultModel comparisonResultModels = await remoteDataSource
            .compareProducts(productModels);
        final ComparisonResultEntity comparisonEntities = comparisonResultModels.toEntity();
        return Right(comparisonEntities);
      } on ServerException catch(e) {
        return Left(ServerFailure(e.message));
      }
    } else {
      return Left(NetworkFailure());
    }
  }

  @override
  Future<Either<Failure, List<ProductEntity>>>
  getProductsForComparison() async {
    try {
      final List<ProductModel> productModels = await localDataSource
          .getProductsForComparison();
      final List<ProductEntity> productEntities = productModels
          .map((productModel) => productModel.toEntity())
          .toList();
      return Right(productEntities);
    } catch (e) {
      return Left(CacheFailure());
    }
  }

  @override
  Future<Either<Failure, void>> removeProductFromCompare(
    String productId,
  ) async {
    try {
      await localDataSource.removeProductFromComparison(productId);
      return Right(null);
    } catch (e) {
      return Left(CacheFailure());
    }
  }

  @override
  Future<Either<Failure, void>> clearProducts() async {
    try{
      await localDataSource.clearProducts();
      return Right(null);
    }catch(e){
      return Left(CacheFailure());
    }
  }
}
