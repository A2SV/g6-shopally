import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../repositories/saved_item_repositories.dart';

class RemoveFromCompare {
  final SavedItemsRepository repository;

  RemoveFromCompare(this.repository);

  Future<Either<Failure, void>> call(String id) async {
    return await repository.removeProduct(id);
  }
}
