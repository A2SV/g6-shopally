import 'package:dartz/dartz.dart';


import 'package:shopallymobile/features/comparing/core/errors/failure.dart';
import 'package:shopallymobile/features/comparing/core/usecases/usecases.dart';
import 'package:shopallymobile/features/comparing/domain/repository/repository.dart';

import '../Entity/product_entity.dart';

class GetProductsForComparison extends UseCase<List<ProductEntity>, NoParams> {
  final Repository repository;

  GetProductsForComparison({required this.repository});

  @override
  Future<Either<Failure, List<ProductEntity>>> call(NoParams params) {
    return repository.getProductsForComparison();
  }
}
