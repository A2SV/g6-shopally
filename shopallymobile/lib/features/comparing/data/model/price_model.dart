
import 'package:shopallymobile/features/comparing/domain/Entity/price_entity.dart';

class PriceModel extends PriceEntity {
  const PriceModel(super.etb, super.usd, super.fxTimestamp);
  factory PriceModel.fromJson(Map<String, dynamic> json) {
    return PriceModel(json['etb'].toDouble(), json['usd'].toDouble(), json['fxTimestamp']);
  }

  Map<String, dynamic> toJson() {
    return {'etb': etb, 'usd': usd, 'fxTimestamp': fxTimestamp};
  }

  PriceEntity toEntity(PriceModel price) {
    return PriceEntity(price.etb, price.usd, price.fxTimestamp);
  }

  factory PriceModel.fromEntity(PriceEntity price) {
    return PriceModel(price.etb, price.usd, price.fxTimestamp);
  }
}
