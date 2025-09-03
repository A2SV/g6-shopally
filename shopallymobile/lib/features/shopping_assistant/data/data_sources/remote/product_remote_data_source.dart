import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

import '../../../../../core/constants/api_uri.dart';

abstract class ProductRemoteDataSource {
  Future<List<Map<String, dynamic>>> fetchProducts(String prompt);
}

class ProductRemoteDataSourceImpl implements ProductRemoteDataSource {
  @override
  Future<List<Map<String, dynamic>>> fetchProducts(String prompt) async {
    final Map<String, String> promptParams = {
      'q': prompt,
    };

    final uri = Uri.parse('${ApiUri.baseUrl}${ApiUri.productsEndpoint}')
      .replace(queryParameters: promptParams);

    final response = await http.get(
      uri,
      headers: {
      'X-Device-ID': 'your-device-id',
      'Accept-Language': 'am',
      },
    );
    debugPrint('Request URL: $uri');
    debugPrint('Response Status: ${response.statusCode}');
    debugPrint('Response Body: ${response.body}');

    if (response.statusCode == 200 || response.statusCode == 201) {
      final decoded = json.decode(response.body) as Map<String, dynamic>;
      final data = decoded['data'] as Map<String, dynamic>?;
      final products = (data?['products'] as List?) ?? [];
      return List<Map<String, dynamic>>.from(
        products.map((e) => Map<String, dynamic>.from(e as Map)),
      );
    } else {
      throw Exception('Failed to load products');
    }
  }
}    // Add prompt/query params here

