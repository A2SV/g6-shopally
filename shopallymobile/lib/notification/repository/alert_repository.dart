import 'package:shopallymobile/notification/datasource/remote/alert_api.dart';

class AlertsRepository {
  Future<String> createAlert({
    required String deviceId,
    required String productId,
    required String productTitle,
    required double currentPrice,
  }) async {
    final result = await AlertsApi.createAlert(
      deviceId: deviceId,
      productId: productId,
      productTitle: productTitle,
      currentPrice: currentPrice,
    );
    return result; // alertId
  }

  Future<void> deleteAlert(String alertId) async {
    await AlertsApi.deleteAlert(alertId);
  }
}
