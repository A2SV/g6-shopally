import 'package:path/path.dart';
import 'package:sqflite/sqflite.dart';

class DatabaseHelper {
  static DatabaseHelper? _databaseHelper;
  static Database? database;

  DatabaseHelper._createInstance();

  factory DatabaseHelper() {
    _databaseHelper ??= DatabaseHelper._createInstance();
    return _databaseHelper!;
  }
  Future<Database?> get db async {
    database ??= await initializeDatabase();
    return database;
  }

  // Removed erroneous code block here.
  Future<Database> initializeDatabase() async {
    String path = join(await getDatabasesPath(), 'shopally.db');
    database = await openDatabase(path, version: 1, onCreate: _createDatabase);
    return database!;
  }

  void _createDatabase(Database db, int newVersion) async {
    await db.execute('''
          CREATE TABLE IF NOT EXISTS Saveditems (
            id TEXT PRIMARY KEY,
            title TEXT,
            imageUrl TEXT,
            price REAL,
            minOrder TEXT,
            rating REAL,
            issaved INTEGER DEFAULT 0
          );
          
          CREATE TABLE IF NOT EXISTS compare (
            id TEXT PRIMARY KEY,
            title TEXT,
            imageUrl TEXT,
            price REAL,
            minOrder TEXT,
            rating REAL,
            iscompare INTEGER DEFAULT 0
          );
        ''');
  }

  Future<List<Map<String, dynamic>>> getall() async {
    Database? db = await this.db;
    var result = await db!.query('Saveditems');
    return result.isNotEmpty ? result : [];
  }

  Future<int> insert(Map<String, dynamic> row) async {
    Database? db = await this.db;
    int result = await db!.insert(
      'Saveditems',
      row,
      conflictAlgorithm: ConflictAlgorithm.ignore,
    );
    print('Item inserted: $result');
    print('Item inserted: $row');
    return result;
  }

  Future<bool> idExists(String table, String id) async {
    final db = await this.db;
    final result = await db!.query(table, where: 'id = ?', whereArgs: [id]);
    return result.isNotEmpty;
  }

  Future<int> update(Map<String, dynamic> row) async {
    Database? db = await this.db;
    String id = row['id'];
    return await db!.update(
      'Saveditems',
      row,
      where: 'id = ?',
      whereArgs: [id],
    );
  }

  Future<int> delete(String id) async {
    Database? db = await this.db;
    return await db!.delete('Saveditems', where: 'id = ?', whereArgs: [id]);
  }

  Future<int> addtoCompare(Map<String, dynamic> row) async {
    Database? db = await this.db;
    int result = await db!.insert(
      'compare',
      row,
      conflictAlgorithm: ConflictAlgorithm.ignore,
    );
    return result;
  }

  Future<int> deletefromCompare(String id) async {
    Database? db = await this.db;
    return await db!.delete('compare', where: 'id = ?', whereArgs: [id]);
  }
}
