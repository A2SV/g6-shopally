import 'dart:convert';

import 'package:path/path.dart';
import 'package:sqflite/sqflite.dart';

class DatabaseShopally {
  static Database? _db;
  static final DatabaseShopally instance = DatabaseShopally._constructor();

  DatabaseShopally._constructor();

  final String productTable = 'products';

  Future<Database> get database async {
    if (_db != null) return _db!;
    _db = await getDatabase();
    return _db!;
  }

  Future<Database> getDatabase() async {
    final dbPath = await getDatabasesPath();
    final databasePath = join(dbPath, 'shopally.db');
    final database = await openDatabase(
      databasePath,
      version: 1,
      onCreate: (db, version) {
        db.execute('''
          CREATE TABLE $productTable(
            id TEXT PRIMARY KEY,
            title TEXT,
            imageUrl TEXT,
            aiMatchPercentage REAL,
            price TEXT,              
            productRating REAL,
            deliveryEstimate TEXT,
            description TEXT,
            productSmallImageUrls TEXT,
            numberSold INTEGER,
            summaryBullets TEXT,     
            deepLinkUrl TEXT,
            tax REAL,
            discount REAL
          )
        ''');

        _insertInitialData(db);
      },
    );
    return database;
  }

  Future<void> _insertInitialData(Database db) async {
    await db.insert(productTable, {
      'id': '1005009254291888',
      'title': 'QIALINO 15-Inch Laptop Sleeve: Waterproof, Shockproof, & Stylish',
      'imageUrl': 'https://ae-pic-a1.aliexpress-media.com/kf/Sddc0076ca42d4a32962f8a617910314bn.jpg',
      'aiMatchPercentage': 30.0,
      'price': jsonEncode({
        'etb': 3857.05,
        'usd': 27.31,
        'fxTimestamp': "2025-09-03T13:42:32.572898678Z",
      }),
      'productRating': 0.0,
      'deliveryEstimate': '',
      'description':
      'Protect your valuable 15-inch laptop with the QIALINO waterproof and shockproof sleeve. Crafted with high-quality materials, this carrying case offers superior protection against accidental bumps, scratches, and spills.',
      'productSmallImageUrls': jsonEncode(null),
      'numberSold': 0,
      'summaryBullets': jsonEncode([
        "Waterproof Design: Keeps your laptop safe from spills and rain.",
        "Shockproof Padding: Provides excellent protection against bumps and drops.",
        "Convenient Carrying: Features a sturdy handle and detachable shoulder strap.",
        "Sleek & Stylish: Complements your laptop with a professional look.",
      ]),
      'deepLinkUrl': 'https://www.aliexpress.com/item/1005009254291888.html',
      'tax': 0.0,
      'discount': 20.0,
    });

    await db.insert(productTable, {
      'id': '1005009710981318',
      'title':
      "54-Piece Anime Sticker Set - 'The Hundred Line' - Perfect for Laptops, Cars, and Gifts!",
      'imageUrl': 'https://ae-pic-a1.aliexpress-media.com/kf/S0fac19d78b5c4c818f6c18c7329db5d4b.jpeg',
      'aiMatchPercentage': 30.0,
      'price': jsonEncode({
        'etb': 927.90,
        'usd': 6.57,
        'fxTimestamp': "2025-09-03T13:42:32.636458648Z",
      }),
      'productRating': 0.0,
      'deliveryEstimate': '',
      'description':
      "Express your love for anime with this 54-piece sticker set featuring characters from 'The Hundred Line - Last Defense Academy'.",
      'productSmallImageUrls': jsonEncode(null),
      'numberSold': 1,
      'summaryBullets': jsonEncode([
        "Unleash your inner anime fan with 54 unique stickers!",
        "Transform ordinary items into personalized masterpieces!",
        "Durable and vibrant, perfect for laptops, cars, and more!",
        "Showcase your love for 'The Hundred Line - Last Defense Academy'!",
      ]),
      'deepLinkUrl': 'https://www.aliexpress.com/item/1005009710981318.html',
      'tax': 0.0,
      'discount': 36.0,
    });
  }



  Future<int> insertProduct(Map<String, dynamic> product) async {
    final db = await database;
    return await db.insert(
      productTable,
      product,
      conflictAlgorithm: ConflictAlgorithm.replace,
    );
  }

  Future<List<Map<String, dynamic>>> getProducts() async {
    final db = await database;
    return await db.query(productTable);
  }

  Future<int> deleteProduct(String id) async {
    final db = await database;
    return await db.delete(productTable, where: 'id = ?', whereArgs: [id]);
  }

  Future<int> clearProducts() async {
    final db = await database;
    return await db.delete(productTable);
  }

  Future close() async {
    final db = await database;
    db.close();
  }
}
