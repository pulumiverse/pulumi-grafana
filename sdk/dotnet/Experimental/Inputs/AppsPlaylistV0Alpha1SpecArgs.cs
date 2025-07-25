// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Grafana.Experimental.Inputs
{

    public sealed class AppsPlaylistV0Alpha1SpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The interval of the playlist.
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("items", required: true)]
        private InputList<Inputs.AppsPlaylistV0Alpha1SpecItemArgs>? _items;

        /// <summary>
        /// The items of the playlist.
        /// </summary>
        public InputList<Inputs.AppsPlaylistV0Alpha1SpecItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.AppsPlaylistV0Alpha1SpecItemArgs>());
            set => _items = value;
        }

        /// <summary>
        /// The title of the playlist.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AppsPlaylistV0Alpha1SpecArgs()
        {
        }
        public static new AppsPlaylistV0Alpha1SpecArgs Empty => new AppsPlaylistV0Alpha1SpecArgs();
    }
}
