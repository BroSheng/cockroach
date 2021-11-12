// Code generated by "stringer"; DO NOT EDIT.

package roachpb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NotLeaseHolderErrType-1]
	_ = x[RangeNotFoundErrType-2]
	_ = x[RangeKeyMismatchErrType-3]
	_ = x[ReadWithinUncertaintyIntervalErrType-4]
	_ = x[TransactionAbortedErrType-5]
	_ = x[TransactionPushErrType-6]
	_ = x[TransactionRetryErrType-7]
	_ = x[TransactionStatusErrType-8]
	_ = x[WriteIntentErrType-9]
	_ = x[WriteTooOldErrType-10]
	_ = x[OpRequiresTxnErrType-11]
	_ = x[ConditionFailedErrType-12]
	_ = x[LeaseRejectedErrType-13]
	_ = x[NodeUnavailableErrType-14]
	_ = x[RaftGroupDeletedErrType-16]
	_ = x[ReplicaCorruptionErrType-17]
	_ = x[ReplicaTooOldErrType-18]
	_ = x[AmbiguousResultErrType-26]
	_ = x[StoreNotFoundErrType-27]
	_ = x[TransactionRetryWithProtoRefreshErrType-28]
	_ = x[IntegerOverflowErrType-31]
	_ = x[UnsupportedRequestErrType-32]
	_ = x[BatchTimestampBeforeGCErrType-34]
	_ = x[TxnAlreadyEncounteredErrType-35]
	_ = x[IntentMissingErrType-36]
	_ = x[MergeInProgressErrType-37]
	_ = x[RangeFeedRetryErrType-38]
	_ = x[IndeterminateCommitErrType-39]
	_ = x[InvalidLeaseErrType-40]
	_ = x[OptimisticEvalConflictsErrType-41]
	_ = x[MinTimestampBoundUnsatisfiableErrType-42]
	_ = x[CommunicationErrType-22]
	_ = x[InternalErrType-25]
}

const (
	_ErrorDetailType_name_0 = "NotLeaseHolderErrTypeRangeNotFoundErrTypeRangeKeyMismatchErrTypeReadWithinUncertaintyIntervalErrTypeTransactionAbortedErrTypeTransactionPushErrTypeTransactionRetryErrTypeTransactionStatusErrTypeWriteIntentErrTypeWriteTooOldErrTypeOpRequiresTxnErrTypeConditionFailedErrTypeLeaseRejectedErrTypeNodeUnavailableErrType"
	_ErrorDetailType_name_1 = "RaftGroupDeletedErrTypeReplicaCorruptionErrTypeReplicaTooOldErrType"
	_ErrorDetailType_name_2 = "CommunicationErrType"
	_ErrorDetailType_name_3 = "InternalErrTypeAmbiguousResultErrTypeStoreNotFoundErrTypeTransactionRetryWithProtoRefreshErrType"
	_ErrorDetailType_name_4 = "IntegerOverflowErrTypeUnsupportedRequestErrType"
	_ErrorDetailType_name_5 = "BatchTimestampBeforeGCErrTypeTxnAlreadyEncounteredErrTypeIntentMissingErrTypeMergeInProgressErrTypeRangeFeedRetryErrTypeIndeterminateCommitErrTypeInvalidLeaseErrTypeOptimisticEvalConflictsErrTypeMinTimestampBoundUnsatisfiableErrType"
)

var (
	_ErrorDetailType_index_0 = [...]uint16{0, 21, 41, 64, 100, 125, 147, 170, 194, 212, 230, 250, 272, 292, 314}
	_ErrorDetailType_index_1 = [...]uint8{0, 23, 47, 67}
	_ErrorDetailType_index_3 = [...]uint8{0, 15, 37, 57, 96}
	_ErrorDetailType_index_4 = [...]uint8{0, 22, 47}
	_ErrorDetailType_index_5 = [...]uint8{0, 29, 57, 77, 99, 120, 146, 165, 195, 232}
)

func (i ErrorDetailType) String() string {
	switch {
	case 1 <= i && i <= 14:
		i -= 1
		return _ErrorDetailType_name_0[_ErrorDetailType_index_0[i]:_ErrorDetailType_index_0[i+1]]
	case 16 <= i && i <= 18:
		i -= 16
		return _ErrorDetailType_name_1[_ErrorDetailType_index_1[i]:_ErrorDetailType_index_1[i+1]]
	case i == 22:
		return _ErrorDetailType_name_2
	case 25 <= i && i <= 28:
		i -= 25
		return _ErrorDetailType_name_3[_ErrorDetailType_index_3[i]:_ErrorDetailType_index_3[i+1]]
	case 31 <= i && i <= 32:
		i -= 31
		return _ErrorDetailType_name_4[_ErrorDetailType_index_4[i]:_ErrorDetailType_index_4[i+1]]
	case 34 <= i && i <= 42:
		i -= 34
		return _ErrorDetailType_name_5[_ErrorDetailType_index_5[i]:_ErrorDetailType_index_5[i+1]]
	default:
		return "ErrorDetailType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
