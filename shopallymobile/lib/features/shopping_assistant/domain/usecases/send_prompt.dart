import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../entities/product_entity.dart';
import '../repositories/prompt_repository.dart';

class SendPrompt {
  final ProductRepository repo;

  SendPrompt(this.repo);

  Future<Either<Failure, List<ProductEntity>>> call(String prompt) async {
    return await repo.sendPrompt(prompt);
  }
}