import 'package:dartz/dartz.dart';
import 'package:shopallymobile/core/errors/failure.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';


class AddToCompare {
  final SavedItemsRepository repository;

  AddToCompare(this.repository);

  Future<Either<Failure, void>> call(Product productToCache) {
    return repository.addtoCompare(productToCache);
  }
}