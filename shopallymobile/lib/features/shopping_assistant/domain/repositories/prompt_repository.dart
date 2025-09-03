

import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entities/product_entity.dart';

abstract class ProductRepository {
  Future<Either<Failure, List<ProductEntity>>> sendPrompt(String prompt);
}