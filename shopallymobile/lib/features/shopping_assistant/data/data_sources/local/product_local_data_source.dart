
abstract class ProductLocalDataSource {
  Future<List<Map<String, dynamic>>> fetchProducts();
}

class ProductLocalDataSourceImpl implements ProductLocalDataSource {
  @override
  Future<List<Map<String, dynamic>>> fetchProducts() async {
    // Simulate a local database call
    await Future.delayed(Duration(seconds: 5));
    return [
      {
        'productName': 'Product 1',
        'productDescription': 'Description 1',
        'id': '1',
        'price': 10.0,
        'imageUrl': 'https://example.com/image1.jpg',
        'inStock': true,
        'rating': 4.5,
      },
      {
        'productName': 'Product 2',
        'productDescription': 'Description 2',
        'id': '2',
        'price': 20.0,
        'imageUrl': 'https://example.com/image2.jpg',
        'inStock': false,
        'rating': 4.0,
      },
    ];
  }
}
