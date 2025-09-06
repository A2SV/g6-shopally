import 'package:shopallymobile/features/saveditem/data/models/product_model.dart';

abstract class ProductLocalDataSource {
  // Future<void> saveProduct(ProductModel productToCache);
  Future<List<ProductModel>> getSavedItems();
  Future<void> saveProduct(ProductModel productToCache);
  Future<void> removeProduct(String id);
  Future<void> addtoCompare(ProductModel productToCache);
  Future<void> removefromCompare(String id);
}