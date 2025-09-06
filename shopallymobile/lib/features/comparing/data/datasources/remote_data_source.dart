import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:shopallymobile/features/comparing/core/constants/constants.dart';
import 'package:shopallymobile/features/comparing/core/errors/exception.dart';
import 'package:shopallymobile/features/comparing/data/model/comparison_result_model.dart';
import 'package:shopallymobile/features/comparing/data/model/product_model.dart';
import 'package:shared_preferences/shared_preferences.dart';


abstract class RemoteDataSource {
  Future<ComparisonResultModel> compareProducts(List<ProductModel> products);
}

class RemoteDataSourceImpl extends RemoteDataSource {
  final http.Client client;
  RemoteDataSourceImpl({required this.client});

  @override
  Future<ComparisonResultModel> compareProducts(
    List<ProductModel> products,
  ) async {
    final url = Uri.parse("$baseUrl/compare");

    final body = {"products": products.map((p) => p.toJson()).toList()};
    final jsonBody = jsonEncode(body);

    try {
      // Read saved language preference (defaults to 'en') similar to search flow
      final prefs = await SharedPreferences.getInstance();
      final langCode = (prefs.getString('lang_code') ?? 'en').toLowerCase();
      final acceptLang = langCode == 'am' ? 'am' : 'en';

      final response = await client.post(
        url,
        headers: {
          "Content-Type": "application/json",
          "X-Device-ID": "12324764",
          "Accept-Language": acceptLang,
        },
        body: jsonBody,
      );

      print('📥 Response status: ${response.statusCode}');
      print('📄 Response body: ${response.body}');

      if (response.statusCode == 200 || response.statusCode == 201) {
        final Map<String, dynamic> jsonData = jsonDecode(response.body);
        final data = jsonData['data'];
        if (data == null || data is! Map<String, dynamic>) {
          throw ServerException(message: 'No comparison data returned from server.');
        }
        return ComparisonResultModel.fromJson(data as Map<String, dynamic>);
      } else if (response.statusCode >= 400 && response.statusCode < 500) {
        throw ServerException(message: "Client error: ${response.statusCode}");
      } else if (response.statusCode >= 500) {
        throw ServerException(message: "Server error: ${response.statusCode}");
      } else {
        throw ServerException(
          message: "Unexpected error: ${response.statusCode}",
        );
      }
    } catch (e, stack) {
      print('❌ Error occurred: $e');
      print(stack);
      rethrow;
    }
  }
}
