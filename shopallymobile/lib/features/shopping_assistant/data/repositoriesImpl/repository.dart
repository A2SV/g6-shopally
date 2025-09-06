import 'package:dartz/dartz.dart';
// import 'package:shopally/features/shopping_assistant/domain/usecases/send_prompt.dart';

import '../../../../core/errors/failure.dart';
import '../../domain/repositories/prompt_repository.dart';
import '../../domain/entities/product_entity.dart';
import '../data_sources/local/product_local_data_source.dart';
import '../data_sources/remote/product_remote_data_source.dart';
import '../models/product_model.dart';

class ProductRepositoryImpl implements ProductRepository {
  final ProductRemoteDataSource remoteDataSource;
  final ProductLocalDataSource localDataSource;

  ProductRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  @override
  Future<Either<Failure, List<ProductEntity>>> sendPrompt(String prompt) async {
    try {
      final productsJson = await remoteDataSource.fetchProducts(prompt);
      final products = productsJson
          .map((json) => ProductModel.fromJson(json) as ProductEntity)
          .toList();
      return Right(products);
    } catch (e) {
      return Left(ServerFailure(message: 'Failed to fetch products: $e'));
    }
  }
}
