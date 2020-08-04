package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type classification struct {
	ID     types.ID     `json:"id" valid:"required~required field id missing"`
	Traits types.Traits `json:"traits" valid:"required~required field traits missing"`
}

var _ mappables.Classification = (*classification)(nil)

func (classification classification) GetID() types.ID { return classification.ID }

func (classification classification) GetTraits() types.Traits { return classification.Traits }

//TODO
func (classification classification) String() string { return "" }

func (classification classification) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(classification)
}
func (classification classification) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &classification)
	return classification
}

func classificationPrototype() traits.Mappable {
	return classification{}
}

func NewClassification(identityID types.ID, traits types.Traits) mappables.Classification {
	return classification{
		ID:     identityID,
		Traits: traits,
	}
}
