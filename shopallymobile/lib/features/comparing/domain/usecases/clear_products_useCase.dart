

import 'package:dartz/dartz.dart';
import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/core/usecases/usecases.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';

class ClearProductsUseCase extends UseCase<void, NoParams> {
  final Repository repository;

  ClearProductsUseCase({required this.repository});
  @override
  Future<Either<Failure, void>> call(NoParams params) async {
    return repository.clearProducts();
  }

}
