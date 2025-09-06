

import 'package:equatable/equatable.dart';

class PriceEntity extends Equatable {
  final double etb;
  final double usd;
  final String fxTimestamp;

  const PriceEntity(this.etb, this.usd, this.fxTimestamp);

  @override
  List<Object?> get props => [etb, usd, fxTimestamp];

}