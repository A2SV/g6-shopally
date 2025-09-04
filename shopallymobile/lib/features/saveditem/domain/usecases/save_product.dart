import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../entities/product.dart';
import '../repositories/saved_item_repositories.dart';

class SaveProduct {
  final SavedItemsRepository repository;

  SaveProduct(this.repository);

  Future<Either<Failure, void>> call(Product productToCache) async {
    return await repository.saveProduct(productToCache);
  }
}
