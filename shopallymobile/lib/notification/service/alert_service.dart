import 'package:shopallymobile/notification/repository/alert_repository.dart';

import '../datasource/remote/fcm_handler.dart';

class AlertsService {
  final AlertsRepository _repo = AlertsRepository();

  // Turn ON alert for this device
  Future<String> enableAlert({
    required String productId,
    required String productTitle,
    required double currentPrice,
  }) async {
    final token = await FcmHandler.getToken();
    if (token == null) throw Exception("FCM token not available");

    return await _repo.createAlert(
      deviceId: token,
      productId: productId,
      productTitle: productTitle,
      currentPrice: currentPrice,
    );
  }

  // Turn OFF alert for this device
  Future<void> disableAlert(String alertId) async {
    await _repo.deleteAlert(alertId);
  }
}
