package main

import (
	"github.com/graphql-go/graphql"
)

var (
	// Feed is the universal Feed type that atom.Feed and rss.Feed gets translated to. It represents a web feed.
	Feed *graphql.Object

	// Item is the universal Item type that atom.Entry and rss.Item gets translated to.  It represents a single entry in a given feed.
	Item *graphql.Object

	// Enclosure is a file associated with a given Item.
	Enclosure *graphql.Object

	// ITunesOwner is the owner of a particular itunes feed.
	ItunesOwner *graphql.Object

	// ITunesCategory is a category element for itunes feeds.
	ItunesCategory *graphql.Object

	// ITunesFeedExtension is a set of extension fields for RSS feeds.
	ItunesFeedExtenstion *graphql.Object

	// ITunesItemExtension is a set of extension fields for RSS items.
	ItunesItemExtenstion *graphql.Object

	// DublinCoreExtension represents a feed extension for the Dublin Core specification.
	DublinCoreExtension *graphql.Object

	// Person is an individual specified in a feed (e.g. an author)
	Person *graphql.Object

	// Image is an image that is the artwork for a given feed or item.
	Image *graphql.Object
)

func stringListFieldMap(names []string) graphql.Fields {
	fm := graphql.Fields{}

	for _, name := range names {
		fm[name] = &graphql.Field{Type: graphql.NewList(graphql.String)}
	}

	return fm
}

func stringFieldMap(names []string) graphql.Fields {
	fm := graphql.Fields{}

	for _, name := range names {
		fm[name] = &graphql.Field{Type: graphql.String}
	}

	return fm
}

func mergeStringFieldMap(names []string, fm graphql.Fields) graphql.Fields {
	for _, name := range names {
		fm[name] = &graphql.Field{Type: graphql.String}
	}

	return fm
}

func init() {
	Image = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Image",
		Description: "Image is an image that is the artwork for a given feed or item.",
		Fields:      stringFieldMap([]string{"url", "title"}),
	})

	Enclosure = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Enclosure",
		Description: "Enclosure is a file associated with a given Item.",
		Fields:      stringFieldMap([]string{"url", "length", "type"}),
	})

	Person = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Person",
		Description: "Person is an individual specified in a feed (e.g. an author)",
		Fields:      stringFieldMap([]string{"name", "email"}),
	})

	ItunesOwner = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ITunesOwner",
		Description: "ITunesOwner is the owner of a particular itunes feed.",
		Fields:      stringFieldMap([]string{"name", "email"}),
	})

	ItunesCategory = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ITunesCategory",
		Description: "ITunesCategory is a category element for itunes feeds.",
		Fields:      stringFieldMap([]string{"text"}),
	})

	ItunesCategory.AddFieldConfig("subcategory", &graphql.Field{Type: graphql.NewList(ItunesCategory)})

	ItunesFeedExtenstion = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ITunesFeedExtension",
		Description: "ITunesFeedExtension is a set of extension fields for RSS feeds.",
		Fields: mergeStringFieldMap([]string{
			"author",
			"block",
			"explicit",
			"keywords",
			"subtite",
			"summary",
			"image",
			"complete",
			"type",
		},
			graphql.Fields{
				"categories": &graphql.Field{Type: graphql.NewList(ItunesCategory)},
				"owner":      &graphql.Field{Type: ItunesOwner},
			}),
	})

	ItunesItemExtenstion = graphql.NewObject(graphql.ObjectConfig{
		Name:        "ITunesItemExtension",
		Description: "ITunesItemExtension is a set of extension fields for RSS items.",
		Fields: stringFieldMap([]string{
			"author",
			"block",
			"duration",
			"explicit",
			"keywords",
			"subtite",
			"summary",
			"image",
			"isClosedCaptioned",
			"episode",
			"season",
			"order",
			"episodeType",
		}),
	})

	DublinCoreExtension = graphql.NewObject(graphql.ObjectConfig{
		Name:        "DublinCoreExtension",
		Description: "DublinCoreExtension represents a feed extension for the Dublin Core specification.",
		Fields: stringListFieldMap([]string{
			"title",
			"creator",
			"author",
			"subject",
			"description",
			"publisher",
			"contributor",
			"date",
			"type",
			"format",
			"identifier",
			"source",
			"language",
			"relation",
			"coverage",
			"rights",
		}),
	})

	Item = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Item",
		Description: "Item is the universal Item type that atom.Entry and rss.Item gets translated to.  It represents a single entry in a given feed.",
		Fields: mergeStringFieldMap([]string{
			"title",
			"description",
			"link",
			"content",
			"updated",
			"published",
			"guid",
		},
			graphql.Fields{
				"updatedParsed":   &graphql.Field{Type: graphql.DateTime},
				"publishedParsed": &graphql.Field{Type: graphql.DateTime},
				"author":          &graphql.Field{Type: Person},
				"image":           &graphql.Field{Type: graphql.String},
				"categories":      &graphql.Field{Type: graphql.NewList(graphql.String)},
				"enclosures":      &graphql.Field{Type: graphql.NewList(Enclosure)},
				"itunesExt":       &graphql.Field{Type: ItunesItemExtenstion},
				"dcExt":           &graphql.Field{Type: DublinCoreExtension},
			}),
	})

	Feed = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Feed",
		Description: "Feed is the universal Feed type that atom.Feed and rss.Feed gets translated to. It represents a web feed.",

		Fields: mergeStringFieldMap([]string{
			"title",
			"description",
			"link",
			"feedLink",
			"updated",
			"language",
			"image",
			"copyright",
			"generator",
			"feedType",
			"feedVersion",
		},
			graphql.Fields{
				"author":          &graphql.Field{Type: Person},
				"updatedParsed":   &graphql.Field{Type: graphql.DateTime},
				"published":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
				"publishedParsed": &graphql.Field{Type: graphql.DateTime},
				"categories":      &graphql.Field{Type: graphql.NewList(graphql.String)},
				"items":           &graphql.Field{Type: graphql.NewList(Item)},
				"itunesExt":       &graphql.Field{Type: ItunesFeedExtenstion},
				"dcExt":           &graphql.Field{Type: DublinCoreExtension},
			}),
	})
}
