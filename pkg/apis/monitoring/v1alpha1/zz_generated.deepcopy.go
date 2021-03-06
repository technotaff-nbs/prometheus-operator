// +build !ignore_autogenerated

// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfig) DeepCopyInto(out *AlertmanagerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfig.
func (in *AlertmanagerConfig) DeepCopy() *AlertmanagerConfig {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigList) DeepCopyInto(out *AlertmanagerConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*AlertmanagerConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AlertmanagerConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigList.
func (in *AlertmanagerConfigList) DeepCopy() *AlertmanagerConfigList {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigSpec) DeepCopyInto(out *AlertmanagerConfigSpec) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(Route)
		(*in).DeepCopyInto(*out)
	}
	if in.Receivers != nil {
		in, out := &in.Receivers, &out.Receivers
		*out = make([]Receiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InhibitRules != nil {
		in, out := &in.InhibitRules, &out.InhibitRules
		*out = make([]InhibitRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigSpec.
func (in *AlertmanagerConfigSpec) DeepCopy() *AlertmanagerConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPConfig) DeepCopyInto(out *HTTPConfig) {
	*out = *in
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(monitoringv1.BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.BearerTokenSecret != nil {
		in, out := &in.BearerTokenSecret, &out.BearerTokenSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(monitoringv1.SafeTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProxyURL != nil {
		in, out := &in.ProxyURL, &out.ProxyURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPConfig.
func (in *HTTPConfig) DeepCopy() *HTTPConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InhibitRule) DeepCopyInto(out *InhibitRule) {
	*out = *in
	if in.TargetMatch != nil {
		in, out := &in.TargetMatch, &out.TargetMatch
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.SourceMatch != nil {
		in, out := &in.SourceMatch, &out.SourceMatch
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Equal != nil {
		in, out := &in.Equal, &out.Equal
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InhibitRule.
func (in *InhibitRule) DeepCopy() *InhibitRule {
	if in == nil {
		return nil
	}
	out := new(InhibitRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Matcher) DeepCopyInto(out *Matcher) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Matcher.
func (in *Matcher) DeepCopy() *Matcher {
	if in == nil {
		return nil
	}
	out := new(Matcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsGenieConfig) DeepCopyInto(out *OpsGenieConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]OpsGenieConfigDetail, len(*in))
		copy(*out, *in)
	}
	if in.Responders != nil {
		in, out := &in.Responders, &out.Responders
		*out = make([]OpsGenieConfigResponder, len(*in))
		copy(*out, *in)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsGenieConfig.
func (in *OpsGenieConfig) DeepCopy() *OpsGenieConfig {
	if in == nil {
		return nil
	}
	out := new(OpsGenieConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsGenieConfigDetail) DeepCopyInto(out *OpsGenieConfigDetail) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsGenieConfigDetail.
func (in *OpsGenieConfigDetail) DeepCopy() *OpsGenieConfigDetail {
	if in == nil {
		return nil
	}
	out := new(OpsGenieConfigDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsGenieConfigResponder) DeepCopyInto(out *OpsGenieConfigResponder) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsGenieConfigResponder.
func (in *OpsGenieConfigResponder) DeepCopy() *OpsGenieConfigResponder {
	if in == nil {
		return nil
	}
	out := new(OpsGenieConfigResponder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyConfig) DeepCopyInto(out *PagerDutyConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.RoutingKey != nil {
		in, out := &in.RoutingKey, &out.RoutingKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceKey != nil {
		in, out := &in.ServiceKey, &out.ServiceKey
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Client != nil {
		in, out := &in.Client, &out.Client
		*out = new(string)
		**out = **in
	}
	if in.ClientURL != nil {
		in, out := &in.ClientURL, &out.ClientURL
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(string)
		**out = **in
	}
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Component != nil {
		in, out := &in.Component, &out.Component
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]PagerDutyConfigDetail, len(*in))
		copy(*out, *in)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyConfig.
func (in *PagerDutyConfig) DeepCopy() *PagerDutyConfig {
	if in == nil {
		return nil
	}
	out := new(PagerDutyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyConfigDetail) DeepCopyInto(out *PagerDutyConfigDetail) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyConfigDetail.
func (in *PagerDutyConfigDetail) DeepCopy() *PagerDutyConfigDetail {
	if in == nil {
		return nil
	}
	out := new(PagerDutyConfigDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Receiver) DeepCopyInto(out *Receiver) {
	*out = *in
	if in.OpsGenieConfigs != nil {
		in, out := &in.OpsGenieConfigs, &out.OpsGenieConfigs
		*out = make([]OpsGenieConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PagerDutyConfigs != nil {
		in, out := &in.PagerDutyConfigs, &out.PagerDutyConfigs
		*out = make([]PagerDutyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SlackConfigs != nil {
		in, out := &in.SlackConfigs, &out.SlackConfigs
		*out = make([]SlackConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WebhookConfigs != nil {
		in, out := &in.WebhookConfigs, &out.WebhookConfigs
		*out = make([]WebhookConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WeChatConfigs != nil {
		in, out := &in.WeChatConfigs, &out.WeChatConfigs
		*out = make([]WeChatConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Receiver.
func (in *Receiver) DeepCopy() *Receiver {
	if in == nil {
		return nil
	}
	out := new(Receiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.GroupBy != nil {
		in, out := &in.GroupBy, &out.GroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Matchers != nil {
		in, out := &in.Matchers, &out.Matchers
		*out = make([]Matcher, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackAction) DeepCopyInto(out *SlackAction) {
	*out = *in
	if in.ConfirmField != nil {
		in, out := &in.ConfirmField, &out.ConfirmField
		*out = new(SlackConfirmationField)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackAction.
func (in *SlackAction) DeepCopy() *SlackAction {
	if in == nil {
		return nil
	}
	out := new(SlackAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfig) DeepCopyInto(out *SlackConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.TitleLink != nil {
		in, out := &in.TitleLink, &out.TitleLink
		*out = new(string)
		**out = **in
	}
	if in.Pretext != nil {
		in, out := &in.Pretext, &out.Pretext
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]SlackField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShortFields != nil {
		in, out := &in.ShortFields, &out.ShortFields
		*out = new(bool)
		**out = **in
	}
	if in.Footer != nil {
		in, out := &in.Footer, &out.Footer
		*out = new(string)
		**out = **in
	}
	if in.Fallback != nil {
		in, out := &in.Fallback, &out.Fallback
		*out = new(string)
		**out = **in
	}
	if in.CallbackID != nil {
		in, out := &in.CallbackID, &out.CallbackID
		*out = new(string)
		**out = **in
	}
	if in.IconEmoji != nil {
		in, out := &in.IconEmoji, &out.IconEmoji
		*out = new(string)
		**out = **in
	}
	if in.IconURL != nil {
		in, out := &in.IconURL, &out.IconURL
		*out = new(string)
		**out = **in
	}
	if in.ImageURL != nil {
		in, out := &in.ImageURL, &out.ImageURL
		*out = new(string)
		**out = **in
	}
	if in.ThumbURL != nil {
		in, out := &in.ThumbURL, &out.ThumbURL
		*out = new(string)
		**out = **in
	}
	if in.LinkNames != nil {
		in, out := &in.LinkNames, &out.LinkNames
		*out = new(bool)
		**out = **in
	}
	if in.MrkdwnIn != nil {
		in, out := &in.MrkdwnIn, &out.MrkdwnIn
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]SlackAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfig.
func (in *SlackConfig) DeepCopy() *SlackConfig {
	if in == nil {
		return nil
	}
	out := new(SlackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfirmationField) DeepCopyInto(out *SlackConfirmationField) {
	*out = *in
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.OkText != nil {
		in, out := &in.OkText, &out.OkText
		*out = new(string)
		**out = **in
	}
	if in.DismissText != nil {
		in, out := &in.DismissText, &out.DismissText
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfirmationField.
func (in *SlackConfirmationField) DeepCopy() *SlackConfirmationField {
	if in == nil {
		return nil
	}
	out := new(SlackConfirmationField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackField) DeepCopyInto(out *SlackField) {
	*out = *in
	if in.Short != nil {
		in, out := &in.Short, &out.Short
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackField.
func (in *SlackField) DeepCopy() *SlackField {
	if in == nil {
		return nil
	}
	out := new(SlackField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeChatConfig) DeepCopyInto(out *WeChatConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.APISecret != nil {
		in, out := &in.APISecret, &out.APISecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.APIURL != nil {
		in, out := &in.APIURL, &out.APIURL
		*out = new(string)
		**out = **in
	}
	if in.CorpID != nil {
		in, out := &in.CorpID, &out.CorpID
		*out = new(string)
		**out = **in
	}
	if in.AgentID != nil {
		in, out := &in.AgentID, &out.AgentID
		*out = new(string)
		**out = **in
	}
	if in.ToUser != nil {
		in, out := &in.ToUser, &out.ToUser
		*out = new(string)
		**out = **in
	}
	if in.ToParty != nil {
		in, out := &in.ToParty, &out.ToParty
		*out = new(string)
		**out = **in
	}
	if in.ToTag != nil {
		in, out := &in.ToTag, &out.ToTag
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.MessageType != nil {
		in, out := &in.MessageType, &out.MessageType
		*out = new(string)
		**out = **in
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeChatConfig.
func (in *WeChatConfig) DeepCopy() *WeChatConfig {
	if in == nil {
		return nil
	}
	out := new(WeChatConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	if in.SendResolved != nil {
		in, out := &in.SendResolved, &out.SendResolved
		*out = new(bool)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.URLSecret != nil {
		in, out := &in.URLSecret, &out.URLSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = new(HTTPConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxAlerts != nil {
		in, out := &in.MaxAlerts, &out.MaxAlerts
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}
