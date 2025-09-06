import 'package:shopallymobile/core/databasehelper/database_helper.dart';
import 'package:shopallymobile/features/saveditem/data/data_sources/local_data/local_data_source.dart';
import 'package:shopallymobile/features/saveditem/data/models/product_model.dart';
import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';

class ProductLocalDataSourceImpl implements ProductLocalDataSource {
  @override
  Future<List<ProductModel>> getSavedItems() async {
    final data = await DatabaseHelper().getall();
    // Use fromMap instead of fromEntity
    print(data.map((e) => ProductModel.fromMap(e)).toList());
    return data.map((e) => ProductModel.fromMap(e)).toList();
  }

  @override
  Future<void> saveProduct(Product productToCache) async {
    print('product inserted: $productToCache');
    // Check if product already exists before inserting
    final exists = await DatabaseHelper().idExists(
      'Saveditems',
      productToCache.id,
    );
    if (!exists) {
      await DatabaseHelper().insert({
        'id': productToCache.id,
        'title': productToCache.title,
        'imageUrl': productToCache.imageUrl,
        'price': productToCache.price,
        'minOrder': productToCache.minOrder,
        'rating': productToCache.rating,
        'issaved': productToCache.issaved,
      });
    }
  }

  @override
  Future<void> removeProduct(String id) async {
    await DatabaseHelper().delete(id);
  }

  @override
  Future<void> addtoCompare(ProductModel productToCache) async {
    // Check if product already exists in compare before inserting
    final exists = await DatabaseHelper().idExists(
      'compare',
      productToCache.id,
    );
    if (!exists) {
      await DatabaseHelper().addtoCompare({
        'id': productToCache.id,
        'title': productToCache.title,
        'imageUrl': productToCache.imageUrl,
        'price': productToCache.price,
        'minOrder': productToCache.minOrder,
        'rating': productToCache.rating,
        'iscompare': productToCache.iscompare,
      });
    }
  }

  @override
  Future<void> removefromCompare(String id) async {
    await DatabaseHelper().deletefromCompare(id);
  }
}
