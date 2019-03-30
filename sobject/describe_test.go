package sobject

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/g8rswimmer/goforce/session"
)

func Test_describe_Describe(t *testing.T) {
	type fields struct {
		session session.Formatter
	}
	type args struct {
		sobject string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    DesribeValue
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Request Error",
			fields: fields{
				session: &mockMetadataSessionFormatter{
					url: "123://wrong",
				},
			},
			args: args{
				sobject: "someobject",
			},
			want:    DesribeValue{},
			wantErr: true,
		},
		{
			name: "Response HTTP Error",
			fields: fields{
				session: &mockMetadataSessionFormatter{
					url: "https://test.salesforce.com",
					client: mockHTTPClient(func(req *http.Request) *http.Response {

						return &http.Response{
							StatusCode: 500,
							Status:     "Some Status",
							Header:     make(http.Header),
						}
					}),
				},
			},
			args: args{
				sobject: "someobject",
			},
			want:    DesribeValue{},
			wantErr: true,
		},
		{
			name: "Response JSON Error",
			fields: fields{
				session: &mockMetadataSessionFormatter{
					url: "https://test.salesforce.com",
					client: mockHTTPClient(func(req *http.Request) *http.Response {
						resp := `
						{`

						return &http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(strings.NewReader(resp)),
							Header:     make(http.Header),
						}
					}),
				},
			},
			args: args{
				sobject: "someobject",
			},
			want:    DesribeValue{},
			wantErr: true,
		},
		{
			name: "Response Passing",
			fields: fields{
				session: &mockMetadataSessionFormatter{
					url: "https://test.salesforce.com",
					client: mockHTTPClient(func(req *http.Request) *http.Response {
						resp := `
						{
							"actionOverrides": [
								{
									"formFactor": "LARGE",
									"isAvailableInTouch": false,
									"name": "New",
									"pageId": "0Abm00000009E2TCAU",
									"url": null
								}
							],
							"activateable": false,
							"childRelationships": [
								{
									"cascadeDelete": false,
									"childSObject": "AcceptedEventRelation",
									"deprecatedAndHidden": false,
									"field": "RelationId",
									"junctionIdListNames": [],
									"junctionReferenceTo": [],
									"relationshipName": "PersonAcceptedEventRelations",
									"restrictedDelete": false
								}
							],
							"compactLayoutable": true,
							"createable": true,
							"custom": false,
							"customSetting": false,
							"deletable": true,
							"deprecatedAndHidden": false,
							"feedEnabled": true,
							"fields": [
								{
									"aggregatable": true,
									"aiPredictionField": false,
									"autoNumber": false,
									"byteLength": 18,
									"calculated": false,
									"calculatedFormula": null,
									"cascadeDelete": false,
									"caseSensitive": false,
									"compoundFieldName": null,
									"controllerName": null,
									"createable": false,
									"custom": false,
									"defaultValue": null,
									"defaultValueFormula": null,
									"defaultedOnCreate": true,
									"dependentPicklist": false,
									"deprecatedAndHidden": false,
									"digits": 0,
									"displayLocationInDecimal": false,
									"encrypted": false,
									"externalId": false,
									"extraTypeInfo": null,
									"filterable": true,
									"filteredLookupInfo": null,
									"formulaTreatNullNumberAsZero": false,
									"groupable": true,
									"highScaleNumber": false,
									"htmlFormatted": false,
									"idLookup": true,
									"inlineHelpText": null,
									"label": "Account ID",
									"length": 18,
									"mask": null,
									"maskType": null,
									"name": "Id",
									"nameField": false,
									"namePointing": false,
									"nillable": false,
									"permissionable": false,
									"picklistValues": [
										{
											"active": true,
											"defaultValue": false,
											"label": "Hand surgery - plastic",
											"validFor": null,
											"value": "Hand surgery - plastic"
										}
									],
									"polymorphicForeignKey": false,
									"precision": 0,
									"queryByDistance": false,
									"referenceTargetField": null,
									"referenceTo": [],
									"relationshipName": null,
									"relationshipOrder": null,
									"restrictedDelete": false,
									"restrictedPicklist": false,
									"scale": 0,
									"searchPrefilterable": false,
									"soapType": "tns:ID",
									"sortable": true,
									"type": "id",
									"unique": false,
									"updateable": false,
									"writeRequiresMasterRead": false
								}
							],
							"hasSubtypes": true,
							"isSubtype": false,
							"keyPrefix": "001",
							"label": "Account",
							"labelPlural": "Accounts",
							"layoutable": true,
							"listviewable": null,
							"lookupLayoutable": null,
							"mergeable": true,
							"mruEnabled": true,
							"name": "Account",
							"namedLayoutInfos": [],
							"networkScopeFieldName": null,
							"queryable": true,
							"recordTypeInfos": [
								{
									"active": true,
									"available": true,
									"defaultRecordTypeMapping": false,
									"developerName": "DeveloperName",
									"master": false,
									"name": "Some Record Name",
									"recordTypeId": "xxx1234",
									"urls": {
										"layout": "/services/data/v44.0/sobjects/Account/describe/layouts/xxx1234"
									}
								}
							],
							"replicateable": true,
							"retrieveable": true,
							"searchLayoutable": true,
							"searchable": true,
							"supportedScopes": [
								{
									"label": "All accounts",
									"name": "everything"
								}
							],
							"triggerable": true,
							"undeletable": true,
							"updateable": true,
							"urls": {
								"compactLayouts": "/services/data/v44.0/sobjects/Account/describe/compactLayouts",
								"rowTemplate": "/services/data/v44.0/sobjects/Account/{ID}",
								"approvalLayouts": "/services/data/v44.0/sobjects/Account/describe/approvalLayouts",
								"uiDetailTemplate": "https:/my.salesforce.com/{ID}",
								"uiEditTemplate": "https://my.salesforce.com/{ID}/e",
								"defaultValues": "/services/data/v44.0/sobjects/Account/defaultValues?recordTypeId&fields",
								"listviews": "/services/data/v44.0/sobjects/Account/listviews",
								"describe": "/services/data/v44.0/sobjects/Account/describe",
								"uiNewRecord": "https://my.salesforce.com/001/e",
								"quickActions": "/services/data/v44.0/sobjects/Account/quickActions",
								"layouts": "/services/data/v44.0/sobjects/Account/describe/layouts",
								"sobject": "/services/data/v44.0/sobjects/Account"
							}
						}`

						return &http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(strings.NewReader(resp)),
							Header:     make(http.Header),
						}
					}),
				},
			},
			args: args{
				sobject: "someobject",
			},
			want:    DesribeValue{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &describe{
				session: tt.fields.session,
			}
			got, err := d.Describe(tt.args.sobject)
			if (err != nil) != tt.wantErr {
				t.Errorf("describe.Describe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("describe.Describe() = %v, want %v", got, tt.want)
			}
		})
	}
}
