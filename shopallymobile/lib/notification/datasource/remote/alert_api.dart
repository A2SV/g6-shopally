import 'dart:convert';
import 'package:http/http.dart' as http;

class AlertsApi {
  static const String baseUrl =
      "https://g6-shopally-3.onrender.com/api/v1/"; // replace with your backend URL

  // Create alert
  static Future<String> createAlert({
    required String deviceId,
    required String productId,
    required String productTitle,
    required double currentPrice,
  }) async {
    final url = Uri.parse("$baseUrl/alerts");
    final body = jsonEncode({
      "deviceId": deviceId,
      "productId": productId,
      "productTitle": productTitle,
      "currentPrice": currentPrice,
    });

    final response = await http.post(
      url,
      headers: {"Content-Type": "application/json"},
      body: body,
    );

    if (response.statusCode == 201) {
      final decoded = jsonDecode(response.body) as Map<String, dynamic>;
      final data = decoded['data'] as Map<String, dynamic>?;
      final alertId = data != null ? data['alertId'] as String? : null;
      if (alertId == null || alertId.isEmpty) {
        throw Exception("Malformed createAlert response: missing alertId");
      }
      return alertId;
    } else {
      throw Exception("Failed to create alert: ${response.body}");
    }
  }

  // Delete alert
  static Future<void> deleteAlert(String alertId) async {
    final url = Uri.parse("$baseUrl/alerts/$alertId");

    final response = await http.delete(url);

    if (response.statusCode != 200) {
      throw Exception("Failed to delete alert: ${response.body}");
    }
  }
}
