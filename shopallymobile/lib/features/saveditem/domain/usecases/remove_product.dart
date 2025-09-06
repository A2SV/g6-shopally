import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../repositories/saved_item_repositories.dart';

class RemoveProduct {
  final SavedItemsRepository repository;

  RemoveProduct(this.repository);

  Future<Either<Failure, void>> call(String id) async {
    return await repository.removeProduct(id);
  }
}
