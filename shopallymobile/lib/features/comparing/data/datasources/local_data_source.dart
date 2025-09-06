import 'package:shopallymobile/features/comparing/core/errors/exception.dart';
import 'package:shopallymobile/features/comparing/data/model/product_model.dart';

import 'database_shopally.dart';

abstract class LocalDataSource {
  Future<List<ProductModel>> getProductsForComparison();
  Future<void> removeProductFromComparison(String productId);
  Future<void> clearProducts();
}

class LocalDataSourceImpl extends LocalDataSource {
  final DatabaseShopally dbHelper;

  LocalDataSourceImpl({required this.dbHelper});

  @override
  Future<List<ProductModel>> getProductsForComparison() async {
    try {
      final List<Map<String, dynamic>> maps = await dbHelper.getProducts();
      final List<ProductModel> productModels =  maps.map((map) => ProductModel.fromDb(map)).toList();
      return productModels;
    } catch (e) {
      throw CacheException();

    }
  }

  @override
  Future<void> removeProductFromComparison(String productId) async {
    try {
      await dbHelper.deleteProduct(productId);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<void> clearProducts() async {
    try {
      await dbHelper.clearProducts();
      return;
    } catch (e) {
      throw CacheException();
    }
  }
}
