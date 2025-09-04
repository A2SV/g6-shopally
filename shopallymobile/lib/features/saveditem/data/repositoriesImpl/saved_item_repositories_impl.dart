import 'package:dartz/dartz.dart';
import 'package:shopallymobile/core/errors/failure.dart';
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source.dart';
import 'package:shopallymobile/features/saveditem/data/models/product_model.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';

class SavedItemsRepositoryImpl implements SavedItemsRepository {
  final ProductLocalDataSource localDataSource;

  SavedItemsRepositoryImpl({required this.localDataSource});

  @override
  Future<Either<Failure, List<Product>>> getSavedItems() async {
    try {
      final products = await localDataSource.getSavedItems();
      return Right(products);
    } catch (e) {
      return Left(DatabaseFailure(message: e.toString()));
    }
  }

  @override
  Future<Either<Failure, void>> saveProduct(Product product) async {
    try {
      final model = ProductModel.fromEntity(product);
      await localDataSource.saveProduct(model);
      return const Right(null);
    } catch (e) {
      return Left(DatabaseFailure(message: e.toString()));
    }
  }

  @override
  Future<Either<Failure, void>> removeProduct(String id) async {
    try {
      await localDataSource.removeProduct(id);
      return const Right(null);
    } catch (e) {
      return Left(DatabaseFailure(message: e.toString()));
    }
  }
  @override
  Future<Either<Failure, void>> addtoCompare(Product productToCache) async {
    try {
      final model = ProductModel.fromEntity(productToCache);
      await localDataSource.addtoCompare(model);
      return const Right(null);
    } catch (e) {
      return Left(DatabaseFailure(message: e.toString()));
    }
  }
  @override
  Future<Either<Failure, void>> removefromCompare(String id) async {
    try {
      await localDataSource.removefromCompare(id);
      return const Right(null);
    } catch (e) {
      return Left(DatabaseFailure(message: e.toString()));
    }
  }
}
