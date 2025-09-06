import 'package:dartz/dartz.dart';
import 'package:shopallymobile/core/errors/failure.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';
import 'package:shopallymobile/features/saveditem/domain/repositories/saved_item_repositories.dart';


class GetSavedItems {
  final SavedItemsRepository repository;

  GetSavedItems(this.repository);

  Future<Either<Failure, List<Product>>> call() {
    return repository.getSavedItems();
  }
}