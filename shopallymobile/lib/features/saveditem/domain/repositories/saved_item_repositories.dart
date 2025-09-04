import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../entities/product.dart';



abstract class SavedItemsRepository {
  Future<Either<Failure, List<Product>>> getSavedItems();
  Future<Either<Failure, void>> saveProduct(Product productToCache);
  Future<Either<Failure, void>> removeProduct(String id);
  Future<Either<Failure, void>> addtoCompare(Product productToCache);
  Future<Either<Failure, void>> removefromCompare(String id);
}