# Schema Types

<details>
  <summary><strong>Table of Contents</strong></summary>

  * [Query](#query)
  * [Mutation](#mutation)
  * [Objects](#objects)
    * [Calendar](#calendar)
    * [ExpoPushToken](#expopushtoken)
    * [Item](#item)
    * [ItemDetail](#itemdetail)
    * [ShareItem](#shareitem)
    * [User](#user)
  * [Inputs](#inputs)
    * [NewCalendar](#newcalendar)
    * [NewItem](#newitem)
    * [NewItemDetail](#newitemdetail)
  * [Scalars](#scalars)
    * [Boolean](#boolean)
    * [Float](#float)
    * [ID](#id)
    * [Int](#int)
    * [String](#string)

</details>

## Query
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>ShareItem</strong></td>
<td valign="top"><a href="#shareitem">ShareItem</a></td>
<td>

公開アイテムを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">id</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>User</strong></td>
<td valign="top"><a href="#user">User</a></td>
<td>

ユーザーを取得する

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>Calendars</strong></td>
<td valign="top">[<a href="#calendar">Calendar</a>]</td>
<td>

カレンダーを期間で取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">startDate</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">endDate</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>Calendar</strong></td>
<td valign="top"><a href="#calendar">Calendar</a></td>
<td>

カレンダーを取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">date</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>ItemDetail</strong></td>
<td valign="top"><a href="#itemdetail">ItemDetail</a></td>
<td>

スケジュール詳細を取得する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">date</td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">itemId</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">itemDetailID</td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
</tbody>
</table>

## Mutation
<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>CreateCalendar</strong></td>
<td valign="top"><a href="#calendar">Calendar</a>!</td>
<td>

カレンダーを作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">calendar</td>
<td valign="top"><a href="#newcalendar">NewCalendar</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>CreateItemDetail</strong></td>
<td valign="top"><a href="#itemdetail">ItemDetail</a>!</td>
<td>

スケジュール詳細を作成する

</td>
</tr>
<tr>
<td colspan="2" align="right" valign="top">itemDetail</td>
<td valign="top"><a href="#newitemdetail">NewItemDetail</a>!</td>
<td></td>
</tr>
</tbody>
</table>

## Objects

### Calendar

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>public</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

true: パブリック、false: プライベート

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>item</strong></td>
<td valign="top"><a href="#item">Item</a>!</td>
<td>

スケジュール

</td>
</tr>
</tbody>
</table>

### ExpoPushToken

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>uid</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deviceId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

デバイスID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>token</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

トークン

</td>
</tr>
</tbody>
</table>

### Item

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>kind</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

種別

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemDetails</strong></td>
<td valign="top">[<a href="#itemdetail">ItemDetail</a>]</td>
<td>

スケジュール詳細

</td>
</tr>
</tbody>
</table>

### ItemDetail

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>kind</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

種類

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>moveMinutes</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>place</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>url</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

URL

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>memo</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

メモ

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>priority</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

表示順

</td>
</tr>
</tbody>
</table>

### ShareItem

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>item</strong></td>
<td valign="top"><a href="#item">Item</a>!</td>
<td>

スケジュール

</td>
</tr>
</tbody>
</table>

### User

<table>
<thead>
<tr>
<th align="left">Field</th>
<th align="right">Argument</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>id</strong></td>
<td valign="top"><a href="#id">ID</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>uid</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

ユーザーID

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>role</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

役割:(管理権限: admin)

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>expoPushTokens</strong></td>
<td valign="top">[<a href="#expopushtoken">ExpoPushToken</a>]</td>
<td>

PUSH通知設定

</td>
</tr>
</tbody>
</table>

## Inputs

### NewCalendar

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>item</strong></td>
<td valign="top"><a href="#newitem">NewItem</a>!</td>
<td>

スケジュール

</td>
</tr>
</tbody>
</table>

### NewItem

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>kind</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

種類

</td>
</tr>
</tbody>
</table>

### NewItemDetail

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>date</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

日付

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>title</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

タイトル

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>itemId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>kind</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>moveMinutes</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>place</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>url</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>memo</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>priority</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td></td>
</tr>
</tbody>
</table>

## Scalars

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### Float

The `Float` scalar type represents signed double-precision fractional values as specified by [IEEE 754](http://en.wikipedia.org/wiki/IEEE_floating_point).

### ID

The `ID` scalar type represents a unique identifier, often used to refetch an object or as key for a cache. The ID type appears in a JSON response as a String; however, it is not intended to be human-readable. When expected as an input type, any string (such as "4") or integer (such as 4) input value will be accepted as an ID.

### Int

The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.

### String

The `String`scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used by GraphQL to represent free-form human-readable text.

