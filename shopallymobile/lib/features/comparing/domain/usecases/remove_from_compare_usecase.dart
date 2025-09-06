import 'package:dartz/dartz.dart';
import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/core/usecases/usecases.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';

class RemoveFromCompareUseCase extends UseCase<void, String> {
  final Repository repository;

  RemoveFromCompareUseCase({required this.repository});
  @override
  Future<Either<Failure, void>> call(String productId) {
    return repository.removeProductFromCompare(productId);
  }
}
